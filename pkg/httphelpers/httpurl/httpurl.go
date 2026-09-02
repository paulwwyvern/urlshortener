package httpurl

import (
	"context"
)

type CtxKey string

var URL CtxKey = "url"

type Data struct {
	URL string
}

// AddURL возвращает запрос, в котором можно хранить информацию о url
func AddURL(ctx context.Context) context.Context {
	data := &Data{}

	return context.WithValue(ctx, URL, data)
}

func SetURL(ctx context.Context, url string) {
	if data, ok := ctx.Value(URL).(*Data); ok {
		data.URL = url
	}
}

func GetURL(ctx context.Context) (string, bool) {
	if data, ok := ctx.Value(URL).(*Data); ok {
		return data.URL, true
	}
	return "", false
}
