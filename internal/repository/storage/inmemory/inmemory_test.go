package inmemory

import (
	"context"
	"testing"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
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
				"a": {OriginalURL: "A"},
				"b": {OriginalURL: "B"},
			},
			shortUrl: "a",
			want:     "A",
			wantErr:  nil,
		},
		{
			name: "Test #2 Url does not exist",
			shortUrlIndex: map[string]*model.URLFile{
				"a": {OriginalURL: "A"},
				"b": {OriginalURL: "B"},
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
				"a": {ShortURL: "A"},
				"b": {ShortURL: "B"},
			},
			originalUrl: "a",
			want:        "A",
			wantErr:     nil,
		},
		{
			name: "Test #2 Url does not exist",
			originalUrlIndex: map[string]*model.URLFile{
				"a": {ShortURL: "A"},
				"b": {ShortURL: "B"},
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
