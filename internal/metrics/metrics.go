package metrics

import (
	"context"
	"github.com/prometheus/client_golang/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
	"time"
)

var (
	GrpcRequestTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_request_total",
			Help: "Total number of gRPC requests.",
		},
		[]string{"method", "status"},
	)

	GrpcRequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "grpc_request_duration_seconds",
			Help: "Histogram of gRPC request durations.",
		},
		[]string{"method"},
	)
)

func init() {
	prometheus.MustRegister(GrpcRequestTotal, GrpcRequestDuration)
}

func UnaryServerInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		startTime := time.Now()
		resp, err = handler(ctx, req)
		var statusLabel string
		if err != nil {
			st, _ := status.FromError(err)
			statusLabel = st.Code().String()
		} else {
			statusLabel = "OK"
		}
		duration := time.Since(startTime).Seconds()
		GrpcRequestTotal.WithLabelValues(info.FullMethod, statusLabel).Inc()
		GrpcRequestDuration.WithLabelValues(info.FullMethod).Observe(duration)
		return resp, err
	}
}
