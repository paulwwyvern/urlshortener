package service

import (
	"errors"
	"fmt"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
)

// Репа где хранятся ссылки
type UrlRepository interface {
	GetURL(string) (string, error)
	SaveURL(shortUrl string, url string) error
}

// Генератор коротких ссылок
type UrlGenerator interface {
	Generate() string
}

type ShortenerService struct {
	hostname string

	urlRepo UrlRepository
	urlGen  UrlGenerator
}

func NewShortener(hostname string, urlRepo UrlRepository, urlGen UrlGenerator) *ShortenerService {
	return &ShortenerService{
		hostname: hostname,

		urlRepo: urlRepo,
		urlGen:  urlGen,
	}
}

func (s *ShortenerService) GenerateURL(url string) (string, error) {
	shortUrl := s.urlGen.Generate()

	err := s.urlRepo.SaveURL(shortUrl, url)

	for err != nil {
		if !errors.Is(err, errs.ErrShortUrlAlreadyExists) {
			return "", err
		}
		shortUrl = s.urlGen.Generate()
		err = s.urlRepo.SaveURL(shortUrl, url)
	}

	return fmt.Sprintf("%s/%s", s.hostname, shortUrl), nil
}

func (s *ShortenerService) GetURL(shortUrl string) (string, error) {
	url, err := s.urlRepo.GetURL(shortUrl)
	if err != nil {
		return "", err
	}
	return url, nil
}
