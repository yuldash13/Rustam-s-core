package outbox

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/prometheus/client_golang/prometheus"
	"go.uber.org/zap"

	"github.com/project/library/config"
	"github.com/project/library/internal/usecase/repository"
)

var (
	outboxSuccessTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "outbox_success_total",
			Help: "Total number of successfully processed outbox messages",
		},
		[]string{"kind"},
	)

	outboxFailedTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "outbox_failed_total",
			Help: "Total number of failed outbox messages processing attempts",
		},
		[]string{"kind", "reason"},
	)

	outboxMessages = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "outbox_messages",
			Help: "Current outbox messages by status and kind",
		},
		[]string{"status", "kind"},
	)
)

func init() {
	prometheus.MustRegister(outboxMessages, outboxSuccessTotal, outboxFailedTotal)
}

type Settings struct {
	AuthorSendURL string
	BookSendURL   string
	Workers       int
	BatchSize     int
	WaitTime      time.Duration
	InProgressTTL time.Duration
}

func RunWorkers(ctx context.Context, logger *zap.Logger, pool *pgxpool.Pool, s Settings) {
	w := s.Workers
	if w < 1 {
		w = 1
	}
	batch := s.BatchSize
	if batch < 1 {
		batch = 1
	}

	wait := config.ParseOutboxDurationMS(os.Getenv("OUTBOX_WAIT_TIME_MS"))
	if wait <= 0 {
		wait = s.WaitTime
	}
	inProgressTTL := config.ParseOutboxDurationMS(os.Getenv("OUTBOX_IN_PROGRESS_TTL_MS"))
	if inProgressTTL <= 0 {
		inProgressTTL = s.InProgressTTL
	}

	logger.Info("outbox workers started",
		zap.Int("workers", w),
		zap.Duration("wait", wait),
		zap.Duration("in_progress_ttl", inProgressTTL),
		zap.String("wait_env", os.Getenv("OUTBOX_WAIT_TIME_MS")),
		zap.String("in_progress_ttl_env", os.Getenv("OUTBOX_IN_PROGRESS_TTL_MS")),
	)

	store := repository.NewOutboxStore(pool)

	startupCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := store.ReclaimAllInProgress(startupCtx); err != nil {
		logger.Error("outbox startup reclaim", zap.Error(err))
	}

	client := &http.Client{Timeout: 30 * time.Second}

	go func() {
		<-ctx.Done()
		reclaimCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		if err := store.ReclaimAllInProgress(reclaimCtx); err != nil {
			logger.Error("outbox reclaim on shutdown", zap.Error(err))
		}
	}()

	for i := 0; i < w; i++ {
		go func() {
			for {
				select {
				case <-ctx.Done():
					return
				default:
				}

				if err := store.ReclaimStale(ctx, inProgressTTL); err != nil {
					logger.Error("outbox reclaim", zap.Error(err))
				}

				msgs, err := store.ClaimBatch(ctx, batch)
				if err != nil {
					if errors.Is(err, context.Canceled) {
						return
					}
					logger.Error("outbox claim", zap.Error(err))
					time.Sleep(wait)
					continue
				}
				if len(msgs) == 0 {
					time.Sleep(wait)
					continue
				}
				for _, m := range msgs {
					processOne(ctx, logger, store, client, s, inProgressTTL, m)
				}
			}
		}()
	}
}

func processOne(
	ctx context.Context,
	logger *zap.Logger,
	store *repository.OutboxStore,
	client *http.Client,
	s Settings,
	inProgressTTL time.Duration,
	m repository.OutboxMessage,
) {
	deliveryCtx := context.WithoutCancel(ctx)

	status, err := store.Status(deliveryCtx, m.Key)
	if err != nil {
		return
	}
	if status == "SUCCESS" {
		return
	}
	if status != "IN_PROGRESS" {
		return
	}

	if err := store.TouchInProgress(deliveryCtx, m.Key); err != nil {
		logger.Error("outbox touch in progress",
			zap.String("key", m.Key),
			zap.Error(err),
		)
		return
	}

	k := kindLabel(m.Kind)

	var body struct {
		ID string `json:"id"`
	}
	if err := json.Unmarshal(m.Data, &body); err != nil || body.ID == "" {
		outboxFailedTotal.WithLabelValues(k, "outbox_bad_payload").Inc()
		logger.Error("outbox bad payload",
			zap.String("key", m.Key),
			zap.Error(err),
		)
		_ = store.MarkRetry(deliveryCtx, m.Key)
		return
	}

	url := s.AuthorSendURL
	if m.Kind == repository.OutboxKindBook {
		url = s.BookSendURL
	}
	if url == "" {
		_ = store.MarkRetry(deliveryCtx, m.Key)
		outboxFailedTotal.WithLabelValues(k, "empty_url").Inc()
		logger.Warn("outbox url is empty",
			zap.String("handler", "OutboxWorker"),
			zap.Int("outbox_kind", m.Kind),
			zap.String("key", m.Key),
		)
		return
	}

	req, err := http.NewRequestWithContext(deliveryCtx, http.MethodPost, url, bytes.NewBufferString(body.ID))
	if err != nil {
		outboxFailedTotal.WithLabelValues(k, "request_build_error").Inc()
		_ = store.MarkRetry(deliveryCtx, m.Key)
		logger.Warn("outbox request build error",
			zap.String("handler", "OutboxWorker"),
			zap.String("operation", "new_request_with_context"),
			zap.Int("outbox_kind", m.Kind),
			zap.String("key", m.Key),
		)
		return
	}
	req.Header.Set("Content-Type", "text/plain; charset=utf-8")

	if m.Kind == repository.OutboxKindBook {
		logger.Info("outbox book post",
			zap.String("handler", "OutboxWorker"),
			zap.String("operation", "outbox_book_post"),
			zap.Int("outbox_kind", m.Kind),
			zap.String("book_id", body.ID),
			zap.String("key", m.Key),
		)
	}

	if m.Kind == repository.OutboxKindAuthor {
		logger.Info("outbox author post",
			zap.String("handler", "OutboxWorker"),
			zap.String("operation", "outbox_author_post"),
			zap.Int("outbox_kind", m.Kind),
			zap.String("author_id", body.ID),
			zap.String("key", m.Key),
		)
	}

	resp, err := client.Do(req)
	if err != nil {
		outboxFailedTotal.WithLabelValues(k, "http_error").Inc()
		if inProgressTTL <= 0 {
			_ = store.MarkRetry(deliveryCtx, m.Key)
			logger.Warn("outbox http error, mark retry (ttl disabled)",
				zap.String("handler", "OutboxWorker"),
				zap.String("operation", "outbox_do"),
				zap.String("key", m.Key),
				zap.Error(err),
			)
			return
		}
		logger.Warn("outbox http error, leaving in_progress",
			zap.String("handler", "OutboxWorker"),
			zap.String("operation", "outbox_do"),
			zap.String("key", m.Key),
			zap.Error(err),
		)
		return
	}
	defer resp.Body.Close()

	entityField := "author_id"
	if m.Kind == repository.OutboxKindBook {
		entityField = "book_id"
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		outboxFailedTotal.WithLabelValues(k, "bad_http_status").Inc()
		logger.Warn("outbox bad http status",
			zap.String("handler", "OutboxWorker"),
			zap.String("operation", "outbox_do"),
			zap.String("key", m.Key),
			zap.String(entityField, body.ID),
			zap.Int("outbox_kind", m.Kind),
			zap.Int("status", resp.StatusCode),
		)
		_ = store.MarkRetry(deliveryCtx, m.Key)
		return
	}

	if err := store.MarkDelivered(deliveryCtx, m.Key); err != nil {
		outboxFailedTotal.WithLabelValues(k, "db_mark_error").Inc()
		logger.Error("outbox mark delivered",
			zap.Int("status", resp.StatusCode),
			zap.Error(err),
			zap.String("key", m.Key),
		)
		return
	}

	logger.Info("outbox delivered",
		zap.String("handler", "OutboxWorker"),
		zap.String("operation", "outbox_deliver"),
		zap.String("key", m.Key),
		zap.String(entityField, body.ID),
		zap.Int("outbox_kind", m.Kind),
		zap.Int("status", resp.StatusCode),
		zap.String("url", url),
	)

	if err := store.MarkSuccess(deliveryCtx, m.Key); err != nil {
		logger.Error("outbox mark success",
			zap.Error(err),
			zap.String("key", m.Key),
		)
	}
	outboxSuccessTotal.WithLabelValues(k).Inc()
}

func kindLabel(kind int) string {
	switch kind {
	case repository.OutboxKindAuthor:
		return "author"
	case repository.OutboxKindBook:
		return "book"
	default:
		return "unknown"
	}
}

func RunMetricsCollector(ctx context.Context, pool *pgxpool.Pool, interval time.Duration, logger *zap.Logger) {
	store := repository.NewOutboxStore(pool)
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			collectCtx, cancel := context.WithTimeout(ctx, 5*time.Second)
			rows, err := store.CountByStatusAndKind(collectCtx)
			cancel()
			if err != nil {
				logger.Error("outbox metrics collect", zap.Error(err))
				continue
			}
			for _, status := range []string{"created", "in_progress", "success"} {
				for _, kind := range []string{"author", "book"} {
					outboxMessages.WithLabelValues(status, kind).Set(0)
				}
			}
			for _, r := range rows {
				outboxMessages.WithLabelValues(strings.ToLower(r.Status), kindLabel(r.Kind)).Set(float64(r.Count))
			}
		}
	}
}
