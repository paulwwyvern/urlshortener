package inmemory

import (
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
	"sync"
)

type Storage struct {
	mu      sync.RWMutex
	storage map[string]string
}

func NewStorage(logger *zap.Logger) *Storage {
	logger.Info("Initializing in-memory storage")
	return &Storage{
		storage: make(map[string]string),
	}
}

func (s *Storage) GetURL(shortUrl string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.storage[shortUrl]
	if !ok {
		return "", errs.ErrShortUrlNotFound
	}

	return url, nil
}

func (s *Storage) SaveURL(shortUrl string, url string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.storage[shortUrl]
	if ok {
		return errs.ErrShortUrlAlreadyExists
	}
	s.storage[shortUrl] = url
	return nil
}
