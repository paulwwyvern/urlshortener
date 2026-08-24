package logger

import (
	"context"

	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

func UnaryLoggerInterceptor(logger *zap.Logger) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		method := info.FullMethod
		userID := httpuser.GetUserID(ctx)

		resp, err := handler(ctx, req)

		logger.Info("GRPC request",
			zap.String("method", method),
			zap.Int32("user id", userID),
			zap.Error(err),
		)

		return resp, err
	}
}
