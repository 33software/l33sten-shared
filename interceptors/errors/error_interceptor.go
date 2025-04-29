package interceptors

import (
	"context"
	"errors"

	"github.com/33software/l33sten-shared/sharederr"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnrayErrorInterceptor() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, request any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, request)

		return resp, ErrToGRPCStatus(err)
		
}}

func ErrToGRPCStatus (err error) error {
	if err == nil {
		return nil
	}

	if status, ok := status.FromError(err); ok {
		return status.Err()
	}

	switch {
	case errors.Is(err, sharederr.ErrDuplicate):
		return status.Error(codes.AlreadyExists, err.Error())
	case errors.Is(err, sharederr.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())
	default:
		return status.Error(codes.Internal, err.Error())
	}
}
