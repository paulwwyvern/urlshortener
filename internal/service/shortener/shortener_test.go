package shortener

import (
	"context"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/paulwwyvern/urlshortener/internal/service/shortener/workers"
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"go.uber.org/zap"
	"testing"
	"time"
)

//go:generate mockgen -source=shortener.go -destination=mock_shortener.go -package=shortener

func TestShortenerService_GenerateURL_Success(t *testing.T) {
	tests := []struct {
		name    string
		baseUrl string

		userId      int32
		url         string
		genShortUrl string

		want    string
		wantErr error
	}{
		{
			name:        "Test #1 Success",
			baseUrl:     "http://localhost:8080",
			userId:      1,
			url:         "http://example.com",
			genShortUrl: "H3dsKvz9o",

			want:    "http://localhost:8080/H3dsKvz9o",
			wantErr: nil,
		}, {
			name:        "Test #2 Success",
			baseUrl:     "http://127.0.0.1:9090",
			userId:      333,
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
			prepo := workers.NewMockPurgeURLRepository(ctrl)

			gen.EXPECT().Generate().Return(tt.genShortUrl)

			repo.EXPECT().SaveURL(gomock.Any(), tt.userId, tt.genShortUrl, tt.url).Return(nil)

			conf := ShortenerServiceConfig{
				BaseUrl:           tt.baseUrl,
				BatchSize:         10,
				URLRepository:     repo,
				URLGenerator:      gen,
				PurgeWorkersCount: 1,
				PurgeWorkersConfig: workers.PurgeWorkerConfig{
					BatchSize:     10,
					PurgeInterval: 1 * time.Second,
					URLRepository: prepo,
				},
			}

			srv := NewShortener(logger, conf)

			shortUrl, err := srv.GenerateURL(context.Background(), tt.userId, tt.url)

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
		prepo := workers.NewMockPurgeURLRepository(ctrl)

		logger := zap.NewNop()

		gomock.InOrder(
			gen.EXPECT().Generate().Return("H3dsKvz9o"),
			repo.EXPECT().SaveURL(gomock.Any(), int32(1234), "H3dsKvz9o", "http://example.com").Return(errs.ErrShortUrlAlreadyExists),
			gen.EXPECT().Generate().Return("DlOi82Xkf"),
			repo.EXPECT().SaveURL(gomock.Any(), int32(1234), "DlOi82Xkf", "http://example.com").Return(nil),
		)

		conf := ShortenerServiceConfig{
			BaseUrl:           "http://example.com",
			BatchSize:         10,
			URLRepository:     repo,
			URLGenerator:      gen,
			PurgeWorkersCount: 1,
			PurgeWorkersConfig: workers.PurgeWorkerConfig{
				BatchSize:     10,
				PurgeInterval: 1 * time.Second,
				URLRepository: prepo,
			},
		}

		srv := NewShortener(logger, conf)

		shortUrl, err := srv.GenerateURL(context.Background(), 1234, "http://example.com")

		assert.ErrorIs(t, err, nil)
		assert.Equal(t, shortUrl, "http://example.com/DlOi82Xkf")
	})
}

func TestShortenerService_GenerateURL_ExistedUrl(t *testing.T) {
	t.Run("Test #1 Exist url", func(t *testing.T) {
		ctrl := gomock.NewController(t)
		defer ctrl.Finish()
		gen := NewMockUrlGenerator(ctrl)
		repo := NewMockUrlRepository(ctrl)
		prepo := workers.NewMockPurgeURLRepository(ctrl)

		logger := zap.NewNop()

		gomock.InOrder(
			gen.EXPECT().Generate().Return("H3dsKvz9o"),
			repo.EXPECT().SaveURL(gomock.Any(), int32(1234), "H3dsKvz9o", "http://example.com").Return(errs.ErrOriginalUrlAlreadyExists),
			repo.EXPECT().GetShortURL(gomock.Any(), "http://example.com").Return("DlOi82Xkf", nil),
		)

		conf := ShortenerServiceConfig{
			BaseUrl:           "http://example.com",
			BatchSize:         10,
			URLRepository:     repo,
			URLGenerator:      gen,
			PurgeWorkersCount: 1,
			PurgeWorkersConfig: workers.PurgeWorkerConfig{
				BatchSize:     10,
				PurgeInterval: 1 * time.Second,
				URLRepository: prepo,
			},
		}

		srv := NewShortener(logger, conf)

		shortUrl, err := srv.GenerateURL(context.Background(), 1234, "http://example.com")

		assert.ErrorIs(t, err, errs.ErrOriginalUrlAlreadyExists)
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
			prepo := workers.NewMockPurgeURLRepository(ctrl)

			repo.EXPECT().GetURL(gomock.Any(), "H3dsKvz9o").Return(tt.want, tt.wantErr)

			conf := ShortenerServiceConfig{
				BaseUrl:           "",
				BatchSize:         10,
				URLRepository:     repo,
				URLGenerator:      gen,
				PurgeWorkersCount: 1,
				PurgeWorkersConfig: workers.PurgeWorkerConfig{
					BatchSize:     10,
					PurgeInterval: 1 * time.Second,
					URLRepository: prepo,
				},
			}

			srv := NewShortener(logger, conf)

			shortUrl, err := srv.GetURL(context.Background(), "H3dsKvz9o")
			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, tt.want, shortUrl)
		})
	}
}

func TestShortenerService_GenerateURLBatch(t *testing.T) {
	type splitBatch struct {
		inRequest  bool
		inResponse bool
		batch      []model.URL
		err        error
	}

	tests := []struct {
		name      string
		userId    int32
		batchSize int
		batch     []splitBatch
		wantErr   error
	}{
		{
			name:      "Test #1 Success",
			userId:    1,
			batchSize: 10,
			batch: []splitBatch{
				{
					inRequest:  true,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "a",
							OriginalURL: "http://example.com",
							ShortURL:    "H3dsKvz9o",
						},
						{
							ID:          "b",
							OriginalURL: "http://yandex.ru",
							ShortURL:    "DlOi82Xkf",
						},
						{
							ID:          "c",
							OriginalURL: "http://google.com",
							ShortURL:    "A34lafApH",
						},
					},
					err: nil,
				},
			},
			wantErr: nil,
		},
		{
			name:      "Test #2 Success split on batches",
			userId:    12,
			batchSize: 2,
			batch: []splitBatch{
				{
					inRequest:  true,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "a",
							OriginalURL: "http://example.com",
							ShortURL:    "H3dsKvz9o",
						},
						{
							ID:          "b",
							OriginalURL: "http://yandex.ru",
							ShortURL:    "DlOi82Xkf",
						},
					},
					err: nil,
				},
				{
					inRequest:  true,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "c",
							OriginalURL: "http://google.com",
							ShortURL:    "A34lafApH",
						},
						{
							ID:          "d",
							OriginalURL: "http://a.com",
							ShortURL:    "H3dsKvz9o",
						},
					},
					err: nil,
				},
			},
			wantErr: nil,
		},
		{
			name:      "Test #3 Success split on batches",
			userId:    124,
			batchSize: 2,
			batch: []splitBatch{
				{
					inRequest:  true,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "a",
							OriginalURL: "http://example.com",
							ShortURL:    "H3dsKvz9o",
						},
						{
							ID:          "b",
							OriginalURL: "http://yandex.ru",
							ShortURL:    "DlOi82Xkf",
						},
					},
					err: nil,
				},
				{
					inRequest:  true,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "c",
							OriginalURL: "http://google.com",
							ShortURL:    "A34lafApH",
						},
					},
					err: nil,
				},
			},
			wantErr: nil,
		},

		{
			name:      "Test #4 Collision",
			userId:    1234,
			batchSize: 2,
			batch: []splitBatch{
				{
					inRequest:  true,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "a",
							OriginalURL: "http://example.com",
							ShortURL:    "H3dsKvz9o",
						},
						{
							ID:          "b",
							OriginalURL: "http://yandex.ru",
							ShortURL:    "DlOi82Xkf",
						},
					},
					err: nil,
				},
				{
					inRequest:  true,
					inResponse: false,
					batch: []model.URL{
						{
							ID:          "c",
							OriginalURL: "http://google.com",
							ShortURL:    "XlU8hVdN8",
						},
						{
							ID:          "d",
							OriginalURL: "http://a.com",
							ShortURL:    "fUfDb0ABN",
						},
					},
					err: errs.ErrShortUrlAlreadyExists,
				},
				{
					inRequest:  false,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "c",
							OriginalURL: "http://google.com",
							ShortURL:    "A34lafApH",
						},
						{
							ID:          "d",
							OriginalURL: "http://a.com",
							ShortURL:    "H3dsKvz9o",
						},
					},
					err: nil,
				},
				{
					inRequest:  true,
					inResponse: true,
					batch: []model.URL{
						{
							ID:          "e",
							OriginalURL: "http://b.ru",
							ShortURL:    "DlOi82Xkf",
						},
						{
							ID:          "f",
							OriginalURL: "http://c.com",
							ShortURL:    "A34lafApH",
						},
					},
					err: nil,
				},
			},
			wantErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()
			gen := NewMockUrlGenerator(ctrl)
			repo := NewMockUrlRepository(ctrl)
			prepo := workers.NewMockPurgeURLRepository(ctrl)

			logger := zap.NewNop()

			req := []model.GenerateURLBatchRequest{}
			wantResp := []model.GenerateURLBatchResponse{}

			for _, b := range tt.batch {
				for _, url := range b.batch {
					gen.EXPECT().Generate().Return(url.ShortURL)
					if b.inRequest {
						req = append(req, model.GenerateURLBatchRequest{
							ID:          url.ID,
							OriginalURL: url.OriginalURL,
						})
					}
					if b.inResponse {
						wantResp = append(wantResp, model.GenerateURLBatchResponse{
							ID:       url.ID,
							ShortURL: "http://example.com/" + url.ShortURL,
						})
					}
				}

				repo.EXPECT().SaveURLBatch(gomock.Any(), tt.userId, b.batch).Return(b.err)
			}

			conf := ShortenerServiceConfig{
				BaseUrl:           "http://example.com",
				BatchSize:         tt.batchSize,
				URLRepository:     repo,
				URLGenerator:      gen,
				PurgeWorkersCount: 1,
				PurgeWorkersConfig: workers.PurgeWorkerConfig{
					BatchSize:     10,
					PurgeInterval: 1 * time.Second,
					URLRepository: prepo,
				},
			}

			srv := NewShortener(logger, conf)

			resp, err := srv.GenerateURLBatch(context.Background(), tt.userId, req)

			assert.ErrorIs(t, err, tt.wantErr)
			assert.Equal(t, resp, wantResp)
		})
	}

}
