package httpuser

import (
	"context"
)

type CtxKey string

var User CtxKey = "user"

func SetUserID(ctx context.Context, userID int32) context.Context {
	return context.WithValue(ctx, User, userID)
}

func GetUserID(ctx context.Context) int32 {
	userID, _ := ctx.Value(User).(int32)
	return userID
}
