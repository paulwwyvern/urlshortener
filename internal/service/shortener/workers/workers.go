package workers

import (
	"context"
	"fmt"
	"go.uber.org/zap"
	"strings"
	"sync"
	"time"
)

type PurgeURLRepository interface {
	PurgeURLBatch(ctx context.Context, urls []string) error
}

type PurgeWorker struct {
	logger *zap.Logger

	workerId int

	batchSize int
	batch     []string
	batchMu   sync.Mutex

	ticker *time.Ticker

	inputCh  <-chan string
	outputCh chan string
	errCh    chan error

	doneCh chan struct{}

	urlRepo PurgeURLRepository
}

type PurgeWorkerConfig struct {
	BatchSize     int
	PurgeInterval time.Duration
	URLRepository PurgeURLRepository

	InputChan <-chan string
}

func NewPurgeWorker(logger *zap.Logger, workerId int, config PurgeWorkerConfig) *PurgeWorker {
	w := &PurgeWorker{
		logger: logger,

		workerId: workerId,

		batchSize: config.BatchSize,
		batch:     make([]string, 0, config.BatchSize),
		ticker:    time.NewTicker(config.PurgeInterval),
		inputCh:   config.InputChan,
		doneCh:    make(chan struct{}),
		errCh:     make(chan error, 1),

		urlRepo: config.URLRepository,
	}

	go w.loop()

	return w
}

func (w *PurgeWorker) GetErrCh() <-chan error {
	return w.errCh
}

func (w *PurgeWorker) loop() {
	defer close(w.doneCh)
	defer close(w.errCh)

	for {
		select {
		case p, ok := <-w.inputCh:
			if !ok {
				w.Flush()
				return
			}
			w.batchMu.Lock()
			w.batch = append(w.batch, p)
			w.batchMu.Unlock()

			if len(w.batch) >= w.batchSize {
				w.Flush()
			}
		case <-w.ticker.C:
			w.Flush()
		}
	}
}

func (w *PurgeWorker) Flush() {
	w.batchMu.Lock()
	defer w.batchMu.Unlock()
	if len(w.batch) == 0 {
		return
	}

	err := w.urlRepo.PurgeURLBatch(context.Background(), w.batch)
	if err != nil {
		w.errCh <- fmt.Errorf("PurgeWorker #%d Flush: %w", w.workerId, err)
		return
	}

	w.logger.Info("Purge worker flushed",
		zap.Int("workerId", w.workerId),
		zap.Int("count", len(w.batch)),
		zap.String("urls", strings.Join(w.batch, ",")),
	)

	w.batch = w.batch[:0]
}

func (w *PurgeWorker) Wait() {
	<-w.doneCh
}

type ErrorWorker struct {
	logger *zap.Logger

	wg    sync.WaitGroup
	errCh chan error

	doneCh chan struct{}
}

func NewErrorWorker(logger *zap.Logger, inputChs ...<-chan error) *ErrorWorker {
	wp := &ErrorWorker{
		logger: logger,
		errCh:  make(chan error, 1),
		doneCh: make(chan struct{}),
	}

	for _, inputCh := range inputChs {
		wp.wg.Add(1)
		go func(ch <-chan error) {
			defer wp.wg.Done()
			for err := range ch {
				wp.errCh <- err
			}
		}(inputCh)
	}

	go wp.loop()

	go func() {
		wp.wg.Wait()
		close(wp.errCh)
	}()

	return wp
}

func (w *ErrorWorker) loop() {
	for err := range w.errCh {
		w.logger.Error("Worker Pool error", zap.Error(err))
	}
	close(w.doneCh)
}

func (w *ErrorWorker) Wait() {
	<-w.doneCh
}
