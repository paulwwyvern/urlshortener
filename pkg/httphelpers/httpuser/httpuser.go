package httpuser

import (
	"context"
	"net/http"
)

type CtxKey string

var User CtxKey = "user"

func SetUserID(r *http.Request, userID int32) {
	*r = *r.WithContext(context.WithValue(r.Context(), User, userID))
}

func GetUserID(r *http.Request) int32 {
	userId, _ := r.Context().Value(User).(int32)
	return userId
}
