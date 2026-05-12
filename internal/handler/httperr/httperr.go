package httperr

import (
	"context"
	"net/http"
)

type CtxKey string

var Error CtxKey = "error"

type HTTPHandler func(w http.ResponseWriter, r *http.Request) error

func Adapt(h HTTPHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		err := h(w, r)
		*r = *r.WithContext(context.WithValue(r.Context(), Error, err))

	}
}

func GetError(r *http.Request) error {
	err, _ := r.Context().Value(Error).(error)
	return err
}
