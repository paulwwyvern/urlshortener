package shortener

import (
	"context"
	"errors"
	"fmt"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/paulwwyvern/urlshortener/internal/service/shortener/workers"
	"go.uber.org/zap"
	"sync"
)

// Репа где хранятся ссылки
type UrlRepository interface {
	GetURL(ctx context.Context, shortUrl string) (string, error)
	GetShortURL(ctx context.Context, url string) (string, error)
	GetUserURL(ctx context.Context, userId int32) ([]model.GetUserURLResponse, error)
	SaveURL(ctx context.Context, userId int32, shortUrl string, url string) error
	SaveURLBatch(ctx context.Context, userId int32, urls []model.URL) error
	SoftDeleteURLBatch(ctx context.Context, userId int32, shortUrls []string) error
	Ping(context.Context) error
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

	purgeCh      chan string
	purgeWg      sync.WaitGroup
	purgeWorkers []*workers.PurgeWorker

	doneCh chan struct{}

	errWorker *workers.ErrorWorker
}

type ShortenerServiceConfig struct {
	BaseUrl   string
	BatchSize int

	URLRepository UrlRepository
	URLGenerator  UrlGenerator

	PurgeWorkersCount  int
	PurgeWorkersConfig workers.PurgeWorkerConfig
}

func NewShortener(logger *zap.Logger, config ShortenerServiceConfig) *ShortenerService {
	logger.Info("Creating shortener service")

	s := &ShortenerService{
		logger: logger,

		baseUrl:   config.BaseUrl,
		batchSize: config.BatchSize,

		urlRepo: config.URLRepository,
		urlGen:  config.URLGenerator,

		purgeCh:      make(chan string, config.PurgeWorkersCount),
		doneCh:       make(chan struct{}),
		purgeWorkers: make([]*workers.PurgeWorker, 0, config.PurgeWorkersCount),
	}

	purgeConfig := config.PurgeWorkersConfig

	// purgeCh - канал по которому приходят урлы для удаления
	// errChs - каналы по которым приходят ошибки от воркеров
	purgeConfig.InputChan = s.purgeCh
	errChs := make([]<-chan error, 0, config.PurgeWorkersCount)

	for i := 0; i < config.PurgeWorkersCount; i++ {
		s.purgeWg.Add(1)

		pw := workers.NewPurgeWorker(logger, i, purgeConfig)
		s.purgeWorkers = append(s.purgeWorkers, pw)

		errChs = append(errChs, pw.GetErrCh())

		go func() {
			pw.Wait()
			s.purgeWg.Done()
		}()
	}

	s.purgeWg.Add(1)
	s.errWorker = workers.NewErrorWorker(logger, errChs...)
	go func() {
		s.errWorker.Wait()
		s.purgeWg.Done()
	}()

	go func() {
		s.purgeWg.Wait()
		close(s.doneCh)
	}()

	return s
}

func (s *ShortenerService) GenerateURL(ctx context.Context, userID int32, url string) (string, error) {

	shortUrl := s.urlGen.Generate()

	err := s.urlRepo.SaveURL(ctx, userID, shortUrl, url)

	var attempts int
	for err != nil {
		if errors.Is(err, errs.ErrOriginalUrlAlreadyExists) {
			shortUrl, err := s.urlRepo.GetShortURL(ctx, url)
			if err != nil {
				return "", err
			}

			return fmt.Sprintf("%s/%s", s.baseUrl, shortUrl), errs.ErrOriginalUrlAlreadyExists

		}
		if !errors.Is(err, errs.ErrShortUrlAlreadyExists) {
			return "", err
		}
		shortUrl = s.urlGen.Generate()
		err = s.urlRepo.SaveURL(ctx, userID, shortUrl, url)
		attempts++
	}

	s.logger.Info("New url generated",
		zap.String("url", url),
		zap.String("shortUrl", shortUrl),
		zap.Int("attempts", attempts),
	)
	return fmt.Sprintf("%s/%s", s.baseUrl, shortUrl), nil
}

func (s *ShortenerService) GenerateURLBatch(ctx context.Context, userID int32, urls []model.GenerateURLBatchRequest) ([]model.GenerateURLBatchResponse, error) {

	batch := make([]model.URL, 0, s.batchSize)
	shortUrls := make([]model.GenerateURLBatchResponse, 0, len(urls))

	var attempts int
	var offset int
	// делим запрос на батчи фикс размера
	for ; offset < len(urls); offset += s.batchSize {
		attempts++

		urlsBatch := urls[offset : offset+min(s.batchSize, len(urls)-offset)]

		// генерим шорты
		for _, url := range urlsBatch {
			var shortUrl string
			shortUrl = s.urlGen.Generate()

			batch = append(batch, model.URL{
				ID:          url.ID,
				ShortURL:    shortUrl,
				OriginalURL: url.OriginalURL,
			})
		}
		// Сам батч меняется, если какие то урлы уже есть в базе
		err := s.urlRepo.SaveURLBatch(ctx, userID, batch)

		// если нашлась коллизия в базе
		if err != nil {
			if errors.Is(err, errs.ErrShortUrlAlreadyExists) {
				// коллизия из за сгенеренных урлов, генерим батч заново
				offset -= s.batchSize
				batch = batch[:0]
				continue
			} else {
				return nil, err
			}
		}

		for _, url := range batch {

			if !url.IsExist {
				s.logger.Info("New url generated",
					zap.String("url", url.OriginalURL),
					zap.String("shortUrl", url.ShortURL),
					zap.Int("attempts", attempts),
				)
			} else {
				s.logger.Info("Trying generated existing url",
					zap.String("url", url.OriginalURL),
					zap.String("shortUrl", url.ShortURL),
				)
			}

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
	return s.urlRepo.GetURL(ctx, shortUrl)
}

func (s *ShortenerService) GetUserURLs(ctx context.Context, userId int32) ([]model.GetUserURLResponse, error) {
	userURL, err := s.urlRepo.GetUserURL(ctx, userId)
	if err != nil {
		return nil, err
	}
	for i := range userURL {
		userURL[i].ShortURL = fmt.Sprintf("%s/%s", s.baseUrl, userURL[i].ShortURL)
	}

	return userURL, nil
}

func (s *ShortenerService) DeleteURLBatch(ctx context.Context, userId int32, shortURLs []string) error {
	err := s.urlRepo.SoftDeleteURLBatch(ctx, userId, shortURLs)
	if err != nil {
		return err
	}

	for _, shortURL := range shortURLs {
		s.purgeCh <- shortURL
	}

	return nil
}

func (s *ShortenerService) Ping(ctx context.Context) error {
	return s.urlRepo.Ping(ctx)
}

func (s *ShortenerService) Close() error {
	close(s.purgeCh)
	<-s.doneCh
	return nil
}
