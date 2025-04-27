package interceptors

import (
	"context"
	"errors"

	"github.com/33software/l33sten-shared/domain"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func UnaryErrorInterceptor() grpc.UnaryServerInterceptor {
	return func (ctx context.Context, request interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		resp, err := handler(ctx, request)

		return resp, ErrorToGRPCStatus(err)
	}
}

func ErrorToGRPCStatus (err error) error {
	if err == nil {
		return nil
	}

	if status, ok := status.FromError(err); ok {
		return status.Err()
	}

	switch {
	case errors.Is(err, domain.ErrNotFound):
		return status.Error(codes.NotFound, err.Error())

		
	default:
		return status.Error(codes.Internal, err.Error())
	}


}