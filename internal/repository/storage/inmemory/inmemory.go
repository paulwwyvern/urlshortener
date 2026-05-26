package inmemory

import (
	"context"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
	"sync"
)

type Storage struct {
	mu sync.RWMutex

	shortUrlIndex    map[string]*model.URLFile
	originalUrlIndex map[string]*model.URLFile
	userIDIndex      map[int32]map[*model.URLFile]struct{}
}

func NewStorage(logger *zap.Logger) (*Storage, error) {
	logger.Info("Initializing in-memory storage")
	return &Storage{
		shortUrlIndex:    make(map[string]*model.URLFile),
		originalUrlIndex: make(map[string]*model.URLFile),
		userIDIndex:      make(map[int32]map[*model.URLFile]struct{}),
	}, nil
}

func (s *Storage) GetAllURLs() map[string]*model.URLFile {
	return s.shortUrlIndex
}

func (s *Storage) GetURL(_ context.Context, shortUrl string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.shortUrlIndex[shortUrl]
	if !ok {
		return "", errs.ErrShortUrlNotFound
	}

	return url.OriginalURL, nil
}

func (s *Storage) GetShortURL(_ context.Context, originalUrl string) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	url, ok := s.originalUrlIndex[originalUrl]
	if !ok {
		return "", errs.ErrOriginalUrlNotFound
	}

	return url.ShortURL, nil
}

func (s *Storage) GetUserURL(_ context.Context, userId int32) ([]model.GetUserURLResponse, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	set := s.userIDIndex[userId]

	userURLs := make([]model.GetUserURLResponse, 0, len(set))
	for url := range set {
		userURLs = append(userURLs, model.GetUserURLResponse{
			OriginalURL: url.OriginalURL,
			ShortURL:    url.ShortURL,
		})
	}

	return userURLs, nil
}

func (s *Storage) SaveURL(_ context.Context, userID int32, shortUrl string, originalUrl string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	_, ok := s.originalUrlIndex[originalUrl]
	if ok {
		return errs.ErrOriginalUrlAlreadyExists
	}
	_, ok = s.shortUrlIndex[shortUrl]
	if ok {
		return errs.ErrShortUrlAlreadyExists
	}

	url := &model.URLFile{
		OriginalURL: originalUrl,
		ShortURL:    shortUrl,
		UserID:      userID,
	}

	s.shortUrlIndex[shortUrl] = url
	s.originalUrlIndex[originalUrl] = url

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
		_, ok := s.shortUrlIndex[url.ShortURL]
		if ok {
			return errs.ErrShortUrlAlreadyExists
		}
		_, ok = s.originalUrlIndex[url.OriginalURL]
		if ok {
			url.ShortURL = s.originalUrlIndex[url.OriginalURL].ShortURL
			url.IsExist = true
			urls[i] = url
		}
	}

	for _, url := range urls {
		if url.IsExist {
			continue
		}

		_, ok := s.originalUrlIndex[url.OriginalURL]
		if ok {
			continue
		}

		saveUrl := &model.URLFile{
			OriginalURL: url.OriginalURL,
			ShortURL:    url.ShortURL,
			UserID:      userID,
		}

		s.shortUrlIndex[url.ShortURL] = saveUrl
		s.originalUrlIndex[url.OriginalURL] = saveUrl

		if s.userIDIndex[userID] == nil {
			s.userIDIndex[userID] = make(map[*model.URLFile]struct{})
		}
		s.userIDIndex[userID][saveUrl] = struct{}{}
	}

	return nil
}

func (s *Storage) SoftDeleteURLBatch(_ context.Context, userId int32, shortUrls []string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	for _, shortUrl := range shortUrls {
		url, ok := s.shortUrlIndex[shortUrl]
		if !ok {
			return errs.ErrShortUrlNotFound
		}

		if url.UserID != userId {
			return errs.ErrShortUrlForbidden
		}
	}

	for _, shortUrl := range shortUrls {
		url := s.shortUrlIndex[shortUrl]
		delete(s.shortUrlIndex, url.ShortURL)
		delete(s.originalUrlIndex, url.OriginalURL)
		if set, ok := s.userIDIndex[userId]; ok {
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
