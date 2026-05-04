package context

import (
	"context"
	"net/http"
)

func WithContext(ctx context.Context) func(http.Handler) http.Handler {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
