package app

import (
	"context"
	"errors"
	"net"
	"net/http"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	stdruntime "runtime"

	"github.com/prometheus/client_golang/prometheus"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracehttp"

	gwruntime "github.com/grpc-ecosystem/grpc-gateway/v2/runtime"

	"github.com/grafana/pyroscope-go"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/sdk/resource"
	"go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.40.0"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"

	"github.com/project/library/config"
	"github.com/project/library/db"
	generated "github.com/project/library/generated/api/library"
	"github.com/project/library/internal/controller"
	"github.com/project/library/internal/usecase/library"
	"github.com/project/library/internal/usecase/outbox"
	"github.com/project/library/internal/usecase/repository"
)

const shutdownTimeout = 10 * time.Second

func Run(logger *zap.Logger, cfg *config.Config) {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	shutdown := initTracer(logger, cfg.Observability.JaegerURL)
	defer func() {
		err := shutdown(ctx)

		if err != nil {
			logger.Error("can not shutdown jaeger collector", zap.Error(err))
		}
	}()

	go runMetricsServer(cfg.Observability.MetricsPort)
	runPyroscope(logger, cfg.Observability.PyroscopeURL)

	dbPool, err := pgxpool.New(ctx, cfg.PG.URL)

	if err != nil {
		logger.Error("can not create pgxpool", zap.Error(err))
		os.Exit(-1)
	}

	defer dbPool.Close()

	db.SetupPostgres(dbPool, logger)

	go repository.RunPostgresMetricsCollector(ctx, dbPool, 15*time.Second, logger)
	
	transactor := repository.NewTransactor(dbPool)
	repo := repository.NewPostgresRepository(dbPool)
	outboxRepo := repository.NewOutboxRepository(dbPool)
	useCases := library.New(logger, repo, repo, transactor, outboxRepo, cfg.Outbox.Enabled)

	ctrl := controller.New(logger, useCases, useCases)

	if cfg.Outbox.Enabled {
		go outbox.RunWorkers(ctx, logger, dbPool, outbox.Settings{
			AuthorSendURL: cfg.Outbox.AuthorSendURL,
			BookSendURL:   cfg.Outbox.BookSendURL,
			Workers:       cfg.Outbox.Workers,
			BatchSize:     cfg.Outbox.BatchSize,
			WaitTime:      config.ParseOutboxDurationMS(os.Getenv("OUTBOX_WAIT_TIME_MS")),
			InProgressTTL: config.ParseOutboxDurationMS(os.Getenv("OUTBOX_IN_PROGRESS_TTL_MS")),
		})
		go outbox.RunMetricsCollector(ctx, dbPool, 15*time.Second, logger)
	}

	grpcSrv, err := newGRPCServer(cfg, logger, ctrl)
	if err != nil {
		logger.Error("can not create grpc server", zap.Error(err))
		os.Exit(-1)
	}
	go grpcSrv.serve(logger)

	gatewaySrv, err := startGatewayServer(cfg, logger)
	if err != nil {
		logger.Error("can not start gateway server", zap.Error(err))
		os.Exit(-1)
	}

	<-ctx.Done()
	logger.Info("shutting down library service")

	shutdownCtx, shutdownCancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer shutdownCancel()

	grpcSrv.stop(shutdownCtx, logger)
	if err := gatewaySrv.Shutdown(shutdownCtx); err != nil {
		logger.Error("gateway shutdown error", zap.Error(err))
	}

	if cfg.Outbox.Enabled {
		outboxStore := repository.NewOutboxStore(dbPool)
		if err := outboxStore.ReclaimAllInProgress(shutdownCtx); err != nil {
			logger.Error("outbox reclaim before shutdown", zap.Error(err))
		}
	}
}

func newGRPCServer(cfg *config.Config, logger *zap.Logger, libraryService generated.LibraryServer) (*grpcServer, error) {
	port := ":" + cfg.GRPC.Port
	lis, err := net.Listen("tcp", port)
	if err != nil {
		return nil, err
	}

	s := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			metricsUnaryInterceptor,
			validateUnaryInterceptor,
			otelgrpc.UnaryServerInterceptor(
				otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
			),
		),
		grpc.ChainStreamInterceptor(
			metricsStreamInterceptor,
			otelgrpc.StreamServerInterceptor(
				otelgrpc.WithTracerProvider(otel.GetTracerProvider()),
			),
		),
	)
	reflection.Register(s)
	generated.RegisterLibraryServer(s, libraryService)

	logger.Info("grpc server listening at port", zap.String("port", port))

	return &grpcServer{srv: s, lis: lis}, nil
}

type grpcServer struct {
	srv *grpc.Server
	lis net.Listener
}

func (g *grpcServer) serve(logger *zap.Logger) {
	if err := g.srv.Serve(g.lis); err != nil && !errors.Is(err, grpc.ErrServerStopped) {
		logger.Error("grpc server listen error", zap.Error(err))
	}
}

func (g *grpcServer) stop(ctx context.Context, logger *zap.Logger) {
	done := make(chan struct{})
	go func() {
		g.srv.GracefulStop()
		close(done)
	}()

	select {
	case <-done:
		logger.Info("grpc server stopped gracefully")
	case <-ctx.Done():
		logger.Warn("grpc graceful shutdown timed out, forcing stop")
		g.srv.Stop()
	}
}

func startGatewayServer(cfg *config.Config, logger *zap.Logger) (*http.Server, error) {
	mux := gwruntime.NewServeMux()
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}

	address := "localhost:" + cfg.GRPC.Port
	if err := generated.RegisterLibraryHandlerFromEndpoint(context.Background(), mux, address, opts); err != nil {
		return nil, err
	}

	gatewayPort := ":" + cfg.GRPC.GatewayPort
	logger.Info("gateway listening at port", zap.String("port", gatewayPort))

	srv := &http.Server{
		Addr:    gatewayPort,
		Handler: mux,
	}
	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logger.Error("gateway listen error", zap.Error(err))
		}
	}()

	return srv, nil
}

func runPyroscope(l *zap.Logger, serverURL string) {
	serverURL = strings.TrimSpace(serverURL)
	if serverURL == "" || strings.EqualFold(serverURL, "off") || strings.EqualFold(serverURL, "disabled") {
		l.Info("pyroscope profiling is disabled")
		return
	}

	stdruntime.SetMutexProfileFraction(1)
	stdruntime.SetBlockProfileRate(1)
	_, err := pyroscope.Start(pyroscope.Config{
		ApplicationName: "leak.app",
		ServerAddress:   serverURL,
		Logger:          pyroscope.StandardLogger,
		ProfileTypes: []pyroscope.ProfileType{
			pyroscope.ProfileCPU,
			pyroscope.ProfileAllocObjects,
			pyroscope.ProfileAllocSpace,
			pyroscope.ProfileInuseObjects,
			pyroscope.ProfileInuseSpace,

			pyroscope.ProfileGoroutines,
			pyroscope.ProfileMutexCount,
			pyroscope.ProfileMutexDuration,
			pyroscope.ProfileBlockCount,
			pyroscope.ProfileBlockDuration,
		},
	})

	if err != nil {
		l.Fatal("can not set up pyroscope", zap.Error(err))
	}
}

func runMetricsServer(port string) {
	http.Handle("/metrics", promhttp.Handler())
	http.ListenAndServe(":"+port, nil)
}

func initTracer(l *zap.Logger, url string) func(ctx context.Context) error {
	exp, err := otlptracehttp.New(
		context.Background(),
		otlptracehttp.WithEndpoint(url), // "jaeger:4318"
		otlptracehttp.WithInsecure(),
	)

	if err != nil {
		l.Fatal("can not create jaeger collector", zap.Error(err))
	}

	tp := trace.NewTracerProvider(
		trace.WithBatcher(exp),
		trace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceName("library-service"),
		)),
	)

	otel.SetTracerProvider(tp)
	return tp.Shutdown
}

var (
	grpcRequestsTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "library_grpc_requests_total",
			Help: "Total gRPC requests",
		},
		[]string{"handler", "code"},
	)

	grpcRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "library_grpc_request_duration_seconds",
			Help:    "gRPC request latency",
			Buckets: prometheus.DefBuckets,
		},
		[]string{"handler"},
	)
)
