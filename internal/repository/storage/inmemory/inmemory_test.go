package inmemory

import (
	"context"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestStorage_GetURL(t *testing.T) {
	tests := []struct {
		name          string
		shortUrlIndex map[string]*model.URLFile
		shortUrl      string
		want          string
		wantErr       error
	}{
		{
			name: "Test #1 Url exists",
			shortUrlIndex: map[string]*model.URLFile{
				"a": &model.URLFile{OriginalURL: "A"},
				"b": &model.URLFile{OriginalURL: "B"},
			},
			shortUrl: "a",
			want:     "A",
			wantErr:  nil,
		},
		{
			name: "Test #2 Url does not exist",
			shortUrlIndex: map[string]*model.URLFile{
				"a": &model.URLFile{OriginalURL: "A"},
				"b": &model.URLFile{OriginalURL: "B"},
			},
			shortUrl: "c",
			want:     "",
			wantErr:  errs.ErrShortUrlNotFound,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewStorage(logger)
			s.shortUrlIndex = tt.shortUrlIndex
			got, err := s.GetURL(context.Background(), tt.shortUrl)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func TestStorage_GetShortURL(t *testing.T) {
	tests := []struct {
		name             string
		originalUrlIndex map[string]*model.URLFile
		originalUrl      string
		want             string
		wantErr          error
	}{
		{
			name: "Test #1 Url exists",
			originalUrlIndex: map[string]*model.URLFile{
				"a": &model.URLFile{ShortURL: "A"},
				"b": &model.URLFile{ShortURL: "B"},
			},
			originalUrl: "a",
			want:        "A",
			wantErr:     nil,
		},
		{
			name: "Test #2 Url does not exist",
			originalUrlIndex: map[string]*model.URLFile{
				"a": &model.URLFile{ShortURL: "A"},
				"b": &model.URLFile{ShortURL: "B"},
			},
			originalUrl: "C",
			want:        "",
			wantErr:     errs.ErrOriginalUrlNotFound,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewStorage(logger)
			s.originalUrlIndex = tt.originalUrlIndex
			got, err := s.GetShortURL(context.Background(), tt.originalUrl)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

/*
func TestStorage_SaveURL(t *testing.T) {
	tests := []struct {
		name             string
		shortUrlIndex    map[string]string
		originalUrlIndex map[string]string

		shortUrl             string
		url                  string
		wantShortUrlIndex    map[string]string
		wantOriginalUrlIndex map[string]string

		wantErr error
	}{
		{
			name: "Test #1 Add new url",

			originalUrlStorage: map[string]string{
				"a": "A",
			},
			shortUrlStorage: map[string]string{
				"A": "a",
			},
			shortUrl: "b",
			url:      "B",
			wantOriginalUrl: map[string]string{
				"a": "A",
				"b": "B",
			},
			wantShortUrl: map[string]string{
				"A": "a",
				"B": "b",
			},
			wantErr: nil,
		}, {
			name: "Test #2 Add existing short url",
			originalUrlStorage: map[string]string{
				"a": "A",
			},
			shortUrlStorage: map[string]string{
				"A": "a",
			},
			shortUrl: "a",
			url:      "B",
			wantOriginalUrl: map[string]string{
				"a": "A",
			},
			wantShortUrl: map[string]string{
				"A": "a",
			},
			wantErr: errs.ErrShortUrlAlreadyExists,
		}, {
			name: "Test #2 Add existing original url",
			originalUrlStorage: map[string]string{
				"a": "A",
			},
			shortUrlStorage: map[string]string{
				"A": "a",
			},
			shortUrl: "b",
			url:      "A",
			wantOriginalUrl: map[string]string{
				"a": "A",
			},
			wantShortUrl: map[string]string{
				"A": "a",
			},
			wantErr: errs.ErrOriginalUrlAlreadyExists,
		},
	}

	logger := zap.NewNop()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewStorage(logger)
			s.originalUrlStorage = tt.originalUrlStorage
			s.shortUrlStorage = tt.shortUrlStorage
			err := s.SaveURL(context.Background(), 1234, tt.shortUrl, tt.url)

			assert.Equal(t, tt.wantOriginalUrl, s.originalUrlStorage)
			assert.Equal(t, tt.wantShortUrl, s.shortUrlStorage)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
*/
