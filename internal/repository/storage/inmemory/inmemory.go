package inmemory

import (
	"context"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
	"sync"
)

type Storage struct {
	mu                 sync.RWMutex
	originalUrlStorage map[string]string
	shortUrlStorage    map[string]string
}

func NewStorage(logger *zap.Logger) (*Storage, error) {
	logger.Info("Initializing in-memory storage")
	return &Storage{
		originalUrlStorage: make(map[string]string),
		shortUrlStorage:    make(map[string]string),
	}, nil
}

func (s *Storage) GetURL(_ context.Context, shortUrl string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.originalUrlStorage[shortUrl]
	if !ok {
		return "", errs.ErrShortUrlNotFound
	}

	return url, nil
}

func (s *Storage) GetShortURL(_ context.Context, url string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	shortUrl, ok := s.shortUrlStorage[url]
	if !ok {
		return "", errs.ErrOriginalUrlNotFound
	}

	return shortUrl, nil
}

func (s *Storage) SaveURL(_ context.Context, shortUrl string, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.shortUrlStorage[url]
	if ok {
		return errs.ErrOriginalUrlAlreadyExists
	}
	_, ok = s.originalUrlStorage[shortUrl]
	if ok {
		return errs.ErrShortUrlAlreadyExists
	}
	s.originalUrlStorage[shortUrl] = url
	s.shortUrlStorage[url] = shortUrl

	return nil
}

func (s *Storage) SaveURLBatch(ctx context.Context, urls []model.URL) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, url := range urls {
		_, ok := s.originalUrlStorage[url.ShortURL]
		if ok {
			return errs.ErrShortUrlAlreadyExists
		}
		_, ok = s.shortUrlStorage[url.OriginalURL]
		if ok {
			url.ShortURL = s.shortUrlStorage[url.OriginalURL]
			url.IsExist = true
			urls[i] = url
		}
	}

	for _, url := range urls {
		_, ok := s.shortUrlStorage[url.OriginalURL]
		if ok {
			continue
		}
		s.originalUrlStorage[url.ShortURL] = url.OriginalURL
		s.shortUrlStorage[url.OriginalURL] = url.ShortURL
	}

	return nil
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) Ping(_ context.Context) error {
	return nil
}
