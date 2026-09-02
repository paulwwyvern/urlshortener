package auth

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"github.com/paulwwyvern/urlshortener/pkg/jwt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAuthFunc(key string) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
		}

		userID, err := jwt.GetUserIDFromJWTToken(key, token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		newCtx := httpuser.SetUserID(ctx, userID)
		return newCtx, nil
	}
}
