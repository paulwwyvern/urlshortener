package service

import (
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"
)

func TestShortenerService_GenerateURL_Success(t *testing.T) {
	tests := []struct {
		name    string
		baseUrl string

		url         string
		genShortUrl string

		want    string
		wantErr error
	}{
		{
			name:        "Test #1 Success",
			baseUrl:     "http://localhost:8080",
			url:         "http://example.com",
			genShortUrl: "H3dsKvz9o",

			want:    "http://localhost:8080/H3dsKvz9o",
			wantErr: nil,
		}, {
			name:        "Test #2 Success",
			baseUrl:     "http://127.0.0.1:9090",
			url:         "http://yandex.ru",
			genShortUrl: "DlOi82Xkf",

			want:    "http://127.0.0.1:9090/DlOi82Xkf",
			wantErr: nil,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			gen := NewMockUrlGenerator(ctrl)
			repo := NewMockUrlRepository(ctrl)

			gen.EXPECT().Generate().Return(tt.genShortUrl)

			repo.EXPECT().SaveURL(tt.genShortUrl, tt.url).Return(nil)

			srv := NewShortener(logger, tt.baseUrl, repo, gen)

			shortUrl, err := srv.GenerateURL(tt.url)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, shortUrl)
		})
	}
}

func TestShortenerService_GenerateURL_Collision(t *testing.T) {
	t.Run("Test #1 Collision", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		gen := NewMockUrlGenerator(ctrl)
		repo := NewMockUrlRepository(ctrl)
		logger := zap.NewNop()

		gomock.InOrder(
			gen.EXPECT().Generate().Return("H3dsKvz9o"),
			repo.EXPECT().SaveURL("H3dsKvz9o", "http://example.com").Return(errs.ErrShortUrlAlreadyExists),
			gen.EXPECT().Generate().Return("DlOi82Xkf"),
			repo.EXPECT().SaveURL("DlOi82Xkf", "http://example.com").Return(nil),
		)

		srv := NewShortener(logger, "http://example.com", repo, gen)

		shortUrl, err := srv.GenerateURL("http://example.com")

		assert.ErrorIs(t, err, nil)
		assert.Equal(t, shortUrl, "http://example.com/DlOi82Xkf")
	})
}

func TestShortenerService_GetURL(t *testing.T) {
	tests := []struct {
		name     string
		shortUrl string
		want     string
		wantErr  error
	}{
		{
			name:     "Test #1 Success",
			shortUrl: "H3dsKvz9o",
			want:     "http://example.com",
			wantErr:  nil,
		}, {
			name:     "Test #2 success",
			shortUrl: "DlOi82Xkf",
			want:     "http://yandex.ru",
			wantErr:  nil,
		}, {
			name:     "Test #3 Not found",
			shortUrl: "DlOi82Xkf",
			want:     "",
			wantErr:  errs.ErrShortUrlNotFound,
		},
	}

	logger := zap.NewNop()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			gen := NewMockUrlGenerator(ctrl)
			repo := NewMockUrlRepository(ctrl)
			repo.EXPECT().GetURL("H3dsKvz9o").Return(tt.want, tt.wantErr)

			srv := NewShortener(logger, "", repo, gen)

			shortUrl, err := srv.GetURL("H3dsKvz9o")
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, shortUrl)
		})
	}
}
