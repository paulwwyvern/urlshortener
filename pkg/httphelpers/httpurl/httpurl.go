package httpurl

import (
	"context"
	"net/http"
)

type CtxKey string

var URL CtxKey = "url"

type Data struct {
	URL string
}

// AddURL возвращает запрос, в котором можно хранить информацию о url
func AddURL(r *http.Request) *http.Request {
	data := &Data{}

	return r.WithContext(context.WithValue(r.Context(), URL, data))
}

func SetURL(r *http.Request, url string) {
	if data, ok := r.Context().Value(URL).(*Data); ok {
		data.URL = url
	}
}

func GetURL(r *http.Request) (string, bool) {
	if data, ok := r.Context().Value(URL).(*Data); ok {
		return data.URL, true
	}
	return "", false
}
