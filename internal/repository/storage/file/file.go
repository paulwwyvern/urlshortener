package file

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/repository/storage/inmemory"
	"go.uber.org/zap"
	"io"
	"os"
)

type Storage struct {
	*inmemory.Storage

	file    *os.File
	encoder *json.Encoder
}

func NewStorage(logger *zap.Logger, file string) (*Storage, error) {
	logger.Info("Initializing in-memory storage with file saving")

	storage, err := readStorage(file)
	if err != nil {
		return nil, err
	}

	f, err := os.OpenFile(file, os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return &Storage{
		Storage: storage,
		file:    f,
		encoder: json.NewEncoder(f),
	}, nil
}

func readStorage(file string) (*inmemory.Storage, error) {
	f, err := os.OpenFile(file, os.O_RDONLY|os.O_CREATE, 0644)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	decoder := json.NewDecoder(f)

	storage, _ := inmemory.NewStorage(zap.NewNop())

	url := model.URLFile{}

	for {
		err = decoder.Decode(&url)
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return nil, err
		}

		err = storage.SaveURL(context.Background(), url.UserID, url.ShortURL, url.OriginalURL)
		if err != nil {
			return nil, err
		}
	}

	return storage, nil

}

func (s *Storage) Close() error {

	urls := s.GetAllURLs()
	for _, url := range urls {
		err := s.encoder.Encode(&url)

		if err != nil {
			return err
		}
	}

	return s.file.Close()
}
