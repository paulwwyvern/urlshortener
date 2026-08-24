package auth

import (
	"context"
	"errors"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/auth"
	"github.com/paulwwyvern/urlshortener/pkg/httphelpers/httpuser"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func GetAuthFunc(key string) func(ctx context.Context) (context.Context, error) {
	return func(ctx context.Context) (context.Context, error) {
		token, err := auth.AuthFromMD(ctx, "bearer")
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
		}

		userID, err := GetUserIDFromJWTToken(key, token)
		if err != nil {
			return nil, status.Error(codes.Unauthenticated, "invalid token")
		}

		newCtx := httpuser.SetUserID(ctx, userID)
		return newCtx, nil
	}
}

type Claims struct {
	jwt.RegisteredClaims
	UserID int32 `json:"user_id"`
}

func GetUserIDFromJWTToken(key string, tokenString string) (int32, error) {
	claims := &Claims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v ", t.Header["alg"])
		}
		return []byte(key), nil
	})

	if err != nil {
		return 0, err
	}

	if !token.Valid {
		return 0, errors.New("invalid token")
	}

	return claims.UserID, nil
}
