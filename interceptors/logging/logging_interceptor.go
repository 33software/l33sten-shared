package interceptors

import (
	"context"
	"log/slog"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

func UnaryLoggingInterceptor(log *slog.Logger) grpc.UnaryServerInterceptor {
	return func (ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		start := time.Now()

		resp, err := handler(ctx, request)

		duration := time.Since(start)

		code := status.Code(err)

		log.Debug("gRPC call results",
		"method", info.FullMethod,
		"result", code.String(),
		"duration", duration.String(),
		"error", err,
	)

	return resp, err
	}
}