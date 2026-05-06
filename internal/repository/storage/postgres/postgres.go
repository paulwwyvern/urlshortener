package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
)

type Storage struct {
	db *sql.DB

	tx map[int64]*sql.Tx
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

	return &Storage{db: db, tx: make(map[int64]*sql.Tx)}, nil
}

func (s *Storage) GetURL(ctx context.Context, shortUrl string) (string, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT url FROM url WHERE short_url = $1`)
	if err != nil {
		return "", errs.ErrInternalError
	}
	defer stmt.Close()

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

func (s *Storage) GetShortURL(ctx context.Context, url string) (string, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT short_url FROM url WHERE url = $1`)
	if err != nil {
		return "", errs.ErrInternalError
	}
	defer stmt.Close()

	var shortUrl string
	err = stmt.QueryRowContext(ctx, url).Scan(&shortUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errs.ErrShortUrlNotFound
		} else {
			return "", errs.ErrInternalError
		}
	}
	return shortUrl, nil
}

func (s *Storage) SaveURL(ctx context.Context, shortUrl string, originalUrl string) error {
	stmt, err := s.db.PrepareContext(ctx, `INSERT INTO url (short_url, url) VALUES ($1, $2)`)
	if err != nil {
		return errs.ErrInternalError
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, shortUrl, originalUrl)
	if err != nil {
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			if pgxErr.Code == pgerrcode.UniqueViolation {
				if pgxErr.ConstraintName == "url_unique" {
					// коллизия по url
					return errs.ErrOriginalUrlAlreadyExists
				} else if pgxErr.ConstraintName == "short_url_unique" {
					// коллизия по short url
					return errs.ErrShortUrlAlreadyExists
				}
			}
		}
		return errs.ErrInternalError
	}
	return nil
}

func (s *Storage) SaveURLBatch(ctx context.Context, urls []model.URL) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return errs.ErrInternalError
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO url (short_url, url) VALUES ($1, $2) 
			ON CONFLICT ON CONSTRAINT url_unique DO UPDATE SET url = EXCLUDED.url 
			RETURNING short_url, (xmax != 0)`)
	if err != nil {
		return errs.ErrInternalError
	}
	defer stmt.Close()

	for i, url := range urls {

		var shortUrl string
		var isExist bool
		err = stmt.QueryRowContext(ctx, url.ShortURL, url.OriginalURL).Scan(&shortUrl, &isExist)
		fmt.Println(err)

		if err != nil {
			var pgxErr *pgconn.PgError
			if errors.As(err, &pgxErr) {
				if pgxErr.Code == pgerrcode.UniqueViolation {
					if pgxErr.ConstraintName == "short_url_unique" {
						// коллизия по short url
						return errs.ErrShortUrlAlreadyExists
					}
				}
			}
			return errs.ErrInternalError
		}

		if isExist {
			url.ShortURL = shortUrl
			url.IsExist = true
			urls[i] = url
			continue
		}
	}
	err = tx.Commit()
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
