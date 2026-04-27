package inmemory

import (
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"
	"testing"
)

func TestStorage_GetURL(t *testing.T) {
	tests := []struct {
		name     string
		storage  map[string]string
		shortUrl string
		want     string
		wantErr  error
	}{
		{
			name: "Test #1 Url exists",
			storage: map[string]string{
				"a": "A",
				"b": "B",
			},
			shortUrl: "a",
			want:     "A",
			wantErr:  nil,
		},
		{
			name: "Test #2 Url does not exist",
			storage: map[string]string{
				"a": "A",
				"b": "B",
			},
			shortUrl: "c",
			want:     "",
			wantErr:  errs.ErrShortUrlNotFound,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorage(logger)
			s.storage = tt.storage
			got, err := s.GetURL(tt.shortUrl)
			assert.Equal(t, tt.want, got)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}

func TestStorage_SaveURL(t *testing.T) {
	tests := []struct {
		name     string
		storage  map[string]string
		shortUrl string
		url      string
		want     map[string]string
		wantErr  error
	}{
		{
			name: "Test #1 Add new url",
			storage: map[string]string{
				"a": "A",
			},
			shortUrl: "b",
			url:      "B",
			want: map[string]string{
				"a": "A",
				"b": "B",
			},
			wantErr: nil,
		}, {
			name: "Test #2 Add existing url",
			storage: map[string]string{
				"a": "A",
			},
			shortUrl: "a",
			url:      "A",
			want: map[string]string{
				"a": "A",
			},
			wantErr: errs.ErrShortUrlAlreadyExists,
		},
	}

	logger := zap.NewNop()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewStorage(logger)
			s.storage = tt.storage
			err := s.SaveURL(tt.shortUrl, tt.url)

			assert.Equal(t, tt.want, s.storage)
			assert.ErrorIs(t, tt.wantErr, err)
		})
	}
}
