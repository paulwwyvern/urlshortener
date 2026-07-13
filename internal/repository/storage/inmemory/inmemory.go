package inmemory

import (
	"context"
	"sync"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/dto"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
)

type Storage struct {
	mu sync.RWMutex

	shortURLIndex    map[string]*model.URLFile
	originalURLIndex map[string]*model.URLFile
	userIDIndex      map[int32]map[*model.URLFile]struct{}
}

func NewStorage(logger *zap.Logger) (*Storage, error) {
	logger.Info("Initializing in-memory storage")
	return &Storage{
		shortURLIndex:    make(map[string]*model.URLFile),
		originalURLIndex: make(map[string]*model.URLFile),
		userIDIndex:      make(map[int32]map[*model.URLFile]struct{}),
	}, nil
}

func (s *Storage) GetAllURLs() map[string]*model.URLFile {
	return s.shortURLIndex
}

func (s *Storage) GetURL(_ context.Context, shortURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.shortURLIndex[shortURL]
	if !ok {
		return "", errs.ErrShortURLNotFound
	}

	return url.OriginalURL, nil
}

func (s *Storage) GetShortURL(_ context.Context, originalURL string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.originalURLIndex[originalURL]
	if !ok {
		return "", errs.ErrOriginalURLNotFound
	}

	return url.ShortURL, nil
}

func (s *Storage) GetUserURL(_ context.Context, userID int32) ([]dto.GetUserURLResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	set := s.userIDIndex[userID]

	userURLs := make([]dto.GetUserURLResponse, 0, len(set))
	for url := range set {
		userURLs = append(userURLs, dto.GetUserURLResponse{
			OriginalURL: url.OriginalURL,
			ShortURL:    url.ShortURL,
		})
	}

	return userURLs, nil
}

func (s *Storage) SaveURL(_ context.Context, userID int32, shortURL string, originalURL string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.originalURLIndex[originalURL]
	if ok {
		return errs.ErrOriginalURLAlreadyExists
	}
	_, ok = s.shortURLIndex[shortURL]
	if ok {
		return errs.ErrShortURLAlreadyExists
	}

	url := &model.URLFile{
		OriginalURL: originalURL,
		ShortURL:    shortURL,
		UserID:      userID,
	}

	s.shortURLIndex[shortURL] = url
	s.originalURLIndex[originalURL] = url

	if s.userIDIndex[userID] == nil {
		s.userIDIndex[userID] = make(map[*model.URLFile]struct{})
	}
	s.userIDIndex[userID][url] = struct{}{}

	return nil
}

func (s *Storage) SaveURLBatch(ctx context.Context, userID int32, urls []model.URL) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for i, url := range urls {
		_, ok := s.shortURLIndex[url.ShortURL]
		if ok {
			return errs.ErrShortURLAlreadyExists
		}
		_, ok = s.originalURLIndex[url.OriginalURL]
		if ok {
			url.ShortURL = s.originalURLIndex[url.OriginalURL].ShortURL
			url.IsExist = true
			urls[i] = url
		}
	}

	for _, url := range urls {
		if url.IsExist {
			continue
		}

		_, ok := s.originalURLIndex[url.OriginalURL]
		if ok {
			continue
		}

		saveURL := &model.URLFile{
			OriginalURL: url.OriginalURL,
			ShortURL:    url.ShortURL,
			UserID:      userID,
		}

		s.shortURLIndex[url.ShortURL] = saveURL
		s.originalURLIndex[url.OriginalURL] = saveURL

		if s.userIDIndex[userID] == nil {
			s.userIDIndex[userID] = make(map[*model.URLFile]struct{})
		}
		s.userIDIndex[userID][saveURL] = struct{}{}
	}

	return nil
}

func (s *Storage) SoftDeleteURLBatch(_ context.Context, userID int32, shortURLs []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, shortURL := range shortURLs {
		url, ok := s.shortURLIndex[shortURL]
		if !ok {
			return errs.ErrShortURLNotFound
		}

		if url.UserID != userID {
			return errs.ErrShortURLForbidden
		}
	}

	for _, shortURL := range shortURLs {
		url := s.shortURLIndex[shortURL]
		delete(s.shortURLIndex, url.ShortURL)
		delete(s.originalURLIndex, url.OriginalURL)
		if set, ok := s.userIDIndex[userID]; ok {
			delete(set, url)
		}
	}

	return nil
}

func (s *Storage) PurgeURLBatch(_ context.Context, _ []string) error {
	return nil
}

func (s *Storage) Close() error {
	return nil
}

func (s *Storage) Ping(_ context.Context) error {
	return nil
}
