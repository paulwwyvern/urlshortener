package httpurl

import (
	"context"
	"net/http"
)

type CtxKey string

var URL CtxKey = "url"

func SetURL(r *http.Request, url string) {
	*r = *r.WithContext(context.WithValue(r.Context(), URL, url))
}

func GetURL(r *http.Request) string {
	url, _ := r.Context().Value(URL).(string)
	return url
}
