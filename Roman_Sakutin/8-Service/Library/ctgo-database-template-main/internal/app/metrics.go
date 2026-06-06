package app

import (
	"context"
	"path"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func init() {
	prometheus.MustRegister(grpcRequestsTotal, grpcRequestDuration)
}

func metricsUnaryInterceptor(
	ctx context.Context,
	req any,
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (any, error) {
	handlerName := path.Base(info.FullMethod)
	start := time.Now()

	resp, err := handler(ctx, req)

	code := status.Code(err).String()
	if err == nil {
		code = codes.OK.String()
	}

	grpcRequestsTotal.WithLabelValues(handlerName, code).Inc()
	grpcRequestDuration.WithLabelValues(handlerName).Observe(time.Since(start).Seconds())

	return resp, err
}

func metricsStreamInterceptor(
	srv any,
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	handlerName := path.Base(info.FullMethod)
	start := time.Now()

	err := handler(srv, ss)

	code := codes.OK.String()
	if err != nil {
		code = status.Code(err).String()
	}

	grpcRequestsTotal.WithLabelValues(handlerName, code).Inc()
	grpcRequestDuration.WithLabelValues(handlerName).Observe(time.Since(start).Seconds())

	return err
}
