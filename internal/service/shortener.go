package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
)

// Репа где хранятся ссылки
type UrlRepository interface {
	GetURL(context.Context, string) (string, error)
	SaveURL(ctx context.Context, shortUrl string, url string) error
}

// Генератор коротких ссылок
type UrlGenerator interface {
	Generate() string
}

type ShortenerService struct {
	logger *zap.Logger

	baseUrl string

	urlRepo UrlRepository
	urlGen  UrlGenerator
}

func NewShortener(logger *zap.Logger, baseUrl string, urlRepo UrlRepository, urlGen UrlGenerator) *ShortenerService {
	logger.Info("Creating shortener service")

	return &ShortenerService{
		logger: logger,

		baseUrl: baseUrl,

		urlRepo: urlRepo,
		urlGen:  urlGen,
	}
}

func (s *ShortenerService) GenerateURL(ctx context.Context, url string) (string, error) {
	shortUrl := s.urlGen.Generate()

	err := s.urlRepo.SaveURL(ctx, shortUrl, url)

	for err != nil {
		if !errors.Is(err, errs.ErrShortUrlAlreadyExists) {
			return "", err
		}
		shortUrl = s.urlGen.Generate()
		err = s.urlRepo.SaveURL(ctx, shortUrl, url)
	}

	s.logger.Info("New url generated",
		zap.String("url", url),
		zap.String("shortUrl", shortUrl),
	)

	return fmt.Sprintf("%s/%s", s.baseUrl, shortUrl), nil
}

func (s *ShortenerService) GetURL(ctx context.Context, shortUrl string) (string, error) {
	url, err := s.urlRepo.GetURL(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}
