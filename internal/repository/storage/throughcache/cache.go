package throughcache

import (
	"context"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/dto"
)

type URLRepository interface {
	GetURL(ctx context.Context, shortURL string) (string, error)
	GetShortURL(ctx context.Context, url string) (string, error)
	GetUserURL(ctx context.Context, userID int32) ([]dto.GetUserURLResponse, error)
	GetURLCount(ctx context.Context) (int, error)
	SaveURL(ctx context.Context, userID int32, shortURL string, url string) error
	SaveURLBatch(ctx context.Context, userID int32, urls []model.URL) error
	SoftDeleteURLBatch(ctx context.Context, userID int32, shortURLs []string) error
	PurgeURLBatch(ctx context.Context, urls []string) error
	CreateUser(ctx context.Context) (int32, error)
	GetUserCount(ctx context.Context) (int, error)
	Ping(context.Context) error
	Close() error
}

type CacheRepository interface {
	Get(string) (string, bool)
	Put(string, string)
	Delete(string)
}

// Cache реализует механизм сквозного кэширования, когда все запросы к репозиторию URLRepository
// проходят через Cache, который пытается достать данные из CacheRepository и если их там нет,
// то подгружает их из URLRepository
type Cache struct {
	cache CacheRepository
	repo  URLRepository
}

func NewCache(cache CacheRepository, repo URLRepository) *Cache {
	return &Cache{
		cache: cache,
		repo:  repo,
	}
}

func (c *Cache) GetURL(ctx context.Context, shortURL string) (string, error) {
	url, ok := c.cache.Get(shortURL)
	if ok {
		return url, nil
	}

	url, err := c.repo.GetURL(ctx, shortURL)
	if err != nil {
		return "", err
	}
	c.cache.Put(shortURL, url)

	return url, nil
}

func (c *Cache) GetShortURL(ctx context.Context, url string) (string, error) {
	return c.repo.GetShortURL(ctx, url)
}

func (c *Cache) GetUserURL(ctx context.Context, userID int32) ([]dto.GetUserURLResponse, error) {
	return c.repo.GetUserURL(ctx, userID)
}

func (c *Cache) GetURLCount(ctx context.Context) (int, error) {
	return c.repo.GetURLCount(ctx)
}

func (c *Cache) SaveURL(ctx context.Context, userID int32, shortURL string, url string) error {
	err := c.repo.SaveURL(ctx, userID, shortURL, url)
	if err != nil {
		return err
	}

	c.cache.Put(shortURL, url)
	return nil
}

func (c *Cache) SaveURLBatch(ctx context.Context, userID int32, urls []model.URL) error {
	err := c.repo.SaveURLBatch(ctx, userID, urls)
	if err != nil {
		return err
	}

	for _, url := range urls {
		c.cache.Put(url.ShortURL, url.OriginalURL)
	}
	return nil
}

func (c *Cache) SoftDeleteURLBatch(ctx context.Context, userID int32, shortURLs []string) error {
	err := c.repo.SoftDeleteURLBatch(ctx, userID, shortURLs)
	if err != nil {
		return err
	}

	for _, shortURL := range shortURLs {
		c.cache.Delete(shortURL)
	}

	return nil
}

func (c *Cache) CreateUser(ctx context.Context) (int32, error) {
	return c.repo.CreateUser(ctx)
}

func (c *Cache) GetUserCount(ctx context.Context) (int, error) {
	return c.repo.GetUserCount(ctx)
}

func (c *Cache) PurgeURLBatch(ctx context.Context, urls []string) error {
	return c.repo.PurgeURLBatch(ctx, urls)
}

func (c *Cache) Ping(ctx context.Context) error {
	return c.repo.Ping(ctx)
}

func (c *Cache) Close() error {
	return c.repo.Close()
}
