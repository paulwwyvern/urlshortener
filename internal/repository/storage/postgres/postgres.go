package postgres

import (
	"context"
	"database/sql"
	"errors"
	"github.com/jackc/pgx/v5"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
)

type Storage struct {
	db *sql.DB
}

func NewStorage(logger *zap.Logger, dsn string) (*Storage, error) {
	logger.Info("Initializing postgres storage")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	logger.Info("Created connection to postgres storage")

	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}

	return &Storage{db: db}, nil

}

func (s *Storage) GetURL(ctx context.Context, shortUrl string) (string, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT url FROM url WHERE short_url = $1`)
	if err != nil {
		return "", errs.ErrInternalError
	}
	var url string
	err = stmt.QueryRowContext(ctx, shortUrl).Scan(&url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errs.ErrShortUrlNotFound
		} else {
			return "", errs.ErrInternalError
		}
	}
	return url, nil
}

func (s *Storage) SaveURL(ctx context.Context, shortUrl string, originalUrl string) error {
	stmt, err := s.db.PrepareContext(ctx, `INSERT INTO url (short_url, url) VALUES ($1, $2)`)
	if err != nil {
		return errs.ErrInternalError
	}

	_, err = stmt.ExecContext(ctx, shortUrl, originalUrl)
	if err != nil {
		return errs.ErrInternalError
	}
	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Ping(ctx context.Context) error {
	return s.db.PingContext(ctx)
}
