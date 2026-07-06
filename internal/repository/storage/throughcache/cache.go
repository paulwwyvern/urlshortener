package throughcache

import (
	"context"

	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/dto"
)

type UrlRepository interface {
	GetURL(ctx context.Context, shortUrl string) (string, error)
	GetShortURL(ctx context.Context, url string) (string, error)
	GetUserURL(ctx context.Context, userId int32) ([]dto.GetUserURLResponse, error)
	SaveURL(ctx context.Context, userId int32, shortUrl string, url string) error
	SaveURLBatch(ctx context.Context, userId int32, urls []model.URL) error
	SoftDeleteURLBatch(ctx context.Context, userId int32, shortUrls []string) error
	PurgeURLBatch(ctx context.Context, urls []string) error
	Ping(context.Context) error
	Close() error
}

type CacheRepository interface {
	Get(string) (string, bool)
	Put(string, string)
	Delete(string)
}

type Cache struct {
	cache CacheRepository
	repo  UrlRepository
}

func NewCache(cache CacheRepository, repo UrlRepository) *Cache {
	return &Cache{
		cache: cache,
		repo:  repo,
	}
}

func (c *Cache) GetURL(ctx context.Context, shortUrl string) (string, error) {
	url, ok := c.cache.Get(shortUrl)
	if ok {
		return url, nil
	}

	url, err := c.repo.GetURL(ctx, shortUrl)
	if err != nil {
		return "", err
	}
	c.cache.Put(shortUrl, url)

	return url, nil
}

func (c *Cache) GetShortURL(ctx context.Context, url string) (string, error) {
	return c.repo.GetShortURL(ctx, url)
}

func (c *Cache) GetUserURL(ctx context.Context, userId int32) ([]dto.GetUserURLResponse, error) {
	return c.repo.GetUserURL(ctx, userId)
}

func (c *Cache) SaveURL(ctx context.Context, userId int32, shortUrl string, url string) error {
	err := c.repo.SaveURL(ctx, userId, shortUrl, url)
	if err != nil {
		return err
	}

	c.cache.Put(shortUrl, url)
	return nil
}

func (c *Cache) SaveURLBatch(ctx context.Context, userId int32, urls []model.URL) error {
	err := c.repo.SaveURLBatch(ctx, userId, urls)
	if err != nil {
		return err
	}

	for _, url := range urls {
		c.cache.Put(url.ShortURL, url.OriginalURL)
	}
	return nil
}

func (c *Cache) SoftDeleteURLBatch(ctx context.Context, userId int32, shortUrls []string) error {
	err := c.repo.SoftDeleteURLBatch(ctx, userId, shortUrls)
	if err != nil {
		return err
	}

	for _, shortUrl := range shortUrls {
		c.cache.Delete(shortUrl)
	}

	return nil
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
