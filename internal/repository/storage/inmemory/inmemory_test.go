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
		shortURLIndex map[string]*model.URLFile
		shortURL      string
		want          string
		wantErr       error
	}{
		{
			name: "Test #1 URL exists",
			shortURLIndex: map[string]*model.URLFile{
				"a": {OriginalURL: "A"},
				"b": {OriginalURL: "B"},
			},
			shortURL: "a",
			want:     "A",
			wantErr:  nil,
		},
		{
			name: "Test #2 URL does not exist",
			shortURLIndex: map[string]*model.URLFile{
				"a": {OriginalURL: "A"},
				"b": {OriginalURL: "B"},
			},
			shortURL: "c",
			want:     "",
			wantErr:  errs.ErrShortURLNotFound,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewStorage(logger)
			s.shortURLIndex = tt.shortURLIndex
			got, err := s.GetURL(context.Background(), tt.shortURL)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func TestStorage_GetShortURL(t *testing.T) {
	tests := []struct {
		name             string
		originalURLIndex map[string]*model.URLFile
		originalURL      string
		want             string
		wantErr          error
	}{
		{
			name: "Test #1 URL exists",
			originalURLIndex: map[string]*model.URLFile{
				"a": {ShortURL: "A"},
				"b": {ShortURL: "B"},
			},
			originalURL: "a",
			want:        "A",
			wantErr:     nil,
		},
		{
			name: "Test #2 URL does not exist",
			originalURLIndex: map[string]*model.URLFile{
				"a": {ShortURL: "A"},
				"b": {ShortURL: "B"},
			},
			originalURL: "C",
			want:        "",
			wantErr:     errs.ErrOriginalURLNotFound,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s, _ := NewStorage(logger)
			s.originalURLIndex = tt.originalURLIndex
			got, err := s.GetShortURL(context.Background(), tt.originalURL)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
