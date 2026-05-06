package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
)

// Репа где хранятся ссылки
type UrlRepository interface {
	GetURL(context.Context, string) (string, error)
	SaveURL(ctx context.Context, shortUrl string, url string) error
	SaveURLBatch(ctx context.Context, urls []model.URL) error
	Ping(context.Context) error
	Close() error
}

// Генератор коротких ссылок
type UrlGenerator interface {
	Generate() string
}

type ShortenerService struct {
	logger *zap.Logger

	baseUrl   string
	batchSize int

	urlRepo UrlRepository
	urlGen  UrlGenerator
}

func NewShortener(logger *zap.Logger, baseUrl string, batchSize int, urlRepo UrlRepository, urlGen UrlGenerator) *ShortenerService {
	logger.Info("Creating shortener service")

	return &ShortenerService{
		logger: logger,

		baseUrl:   baseUrl,
		batchSize: batchSize,

		urlRepo: urlRepo,
		urlGen:  urlGen,
	}
}

func (s *ShortenerService) GenerateURL(ctx context.Context, url string) (string, error) {
	shortUrl := s.urlGen.Generate()

	err := s.urlRepo.SaveURL(ctx, shortUrl, url)

	var attempts int
	for err != nil {
		if !errors.Is(err, errs.ErrShortUrlAlreadyExists) {
			return "", err
		}
		shortUrl = s.urlGen.Generate()
		err = s.urlRepo.SaveURL(ctx, shortUrl, url)
		attempts++
	}

	s.logger.Info("New url generated",
		zap.String("url", url),
		zap.String("shortUrl", shortUrl),
		zap.Int("attempts", attempts),
	)

	return fmt.Sprintf("%s/%s", s.baseUrl, shortUrl), nil
}

func (s *ShortenerService) GenerateURLBatch(ctx context.Context, urls []model.GenerateURLBatchRequest) ([]model.GenerateURLBatchResponse, error) {
	batch := make([]model.URL, 0, s.batchSize)
	shortUrls := make([]model.GenerateURLBatchResponse, 0, len(urls))

	exist := make(map[string]struct{})

	var attempts int
	var offset int
	// делим запрос на батчи фикс размера
	for ; offset < len(urls); offset += s.batchSize {
		attempts++

		urlsBatch := urls[offset : offset+min(s.batchSize, len(urls)-offset)]

		// генерим шорты
		for _, url := range urlsBatch {
			var shortUrl string
			for {
				shortUrl = s.urlGen.Generate()
				if _, ok := exist[shortUrl]; ok {
					continue
				}
				exist[shortUrl] = struct{}{}
				break
			}

			batch = append(batch, model.URL{
				ID:          url.ID,
				ShortURL:    shortUrl,
				OriginalURL: url.OriginalURL,
			})
		}
		// если нашлась коллизия в базе
		if err := s.urlRepo.SaveURLBatch(ctx, batch); err != nil {
			if errors.Is(err, errs.ErrShortUrlAlreadyExists) {
				offset -= s.batchSize
				batch = batch[:0]
				continue
			} else {
				return nil, err
			}
		}

		for _, url := range batch {
			s.logger.Info("New url generated",
				zap.String("url", url.OriginalURL),
				zap.String("shortUrl", url.ShortURL),
				zap.Int("attempts", attempts),
			)

			shortUrls = append(shortUrls, model.GenerateURLBatchResponse{
				ID:       url.ID,
				ShortURL: fmt.Sprintf("%s/%s", s.baseUrl, url.ShortURL),
			})
		}
		attempts = 0
		batch = batch[:0]
	}
	return shortUrls, nil

}

func (s *ShortenerService) GetURL(ctx context.Context, shortUrl string) (string, error) {
	url, err := s.urlRepo.GetURL(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}

func (s *ShortenerService) Ping(ctx context.Context) error {
	return s.urlRepo.Ping(ctx)
}
