package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"go.uber.org/zap"
)

const (
	shortUrlConstraintName    = "short_url_unique"
	originalUrlConstraintName = "url_unique"
)

type Storage struct {
	db *sql.DB

	tx map[int64]*sql.Tx
}

func NewStorage(logger *zap.Logger, dsn string, migrate bool, migrationSource string) (*Storage, error) {
	if migrate {
		logger.Info("Initializing migration")
		err := Migrate(migrationSource, dsn)
		if err != nil {
			return nil, err
		}
		logger.Info("Migration complete")
	}

	logger.Info("Initializing postgres storage")

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}

	err = db.PingContext(context.Background())
	if err != nil {
		return nil, err
	}
	logger.Info("Created connection to postgres storage")

	return &Storage{db: db, tx: make(map[int64]*sql.Tx)}, nil
}

func Migrate(source string, dsn string) error {
	m, err := migrate.New("file://"+source, dsn)
	if err != nil {
		return err
	}
	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}
	return nil
}

func (s *Storage) GetURL(ctx context.Context, shortUrl string) (string, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT url FROM url WHERE short_url = $1`)
	if err != nil {
		return "", fmt.Errorf("GetURL: failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var url string
	err = stmt.QueryRowContext(ctx, shortUrl).Scan(&url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errs.ErrShortUrlNotFound
		} else {
			return "", fmt.Errorf("GetURL: failed to get url: %w", err)
		}
	}
	return url, nil
}

func (s *Storage) GetShortURL(ctx context.Context, url string) (string, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT short_url FROM url WHERE url = $1`)
	if err != nil {
		return "", fmt.Errorf("GetShortURL: failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var shortUrl string
	err = stmt.QueryRowContext(ctx, url).Scan(&shortUrl)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", errs.ErrShortUrlNotFound
		} else {
			return "", fmt.Errorf("GetShortURL: failed to get url: %w", err)
		}
	}
	return shortUrl, nil
}

func (s *Storage) GetUserURL(ctx context.Context, userID int32) ([]model.GetUserURLResponse, error) {
	stmt, err := s.db.PrepareContext(ctx, `SELECT short_url, url FROM url WHERE user_id = $1`)
	if err != nil {
		return nil, fmt.Errorf("GetUserURL: failed to prepare query: %w", err)
	}
	defer stmt.Close()

	var userURL []model.GetUserURLResponse
	rows, err := stmt.QueryContext(ctx, userID)
	if err != nil {
		return nil, fmt.Errorf("GetUserURL: failed to query rows: %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		var shortUrl string
		var url string
		err = rows.Scan(&shortUrl, &url)
		if err != nil {
			return nil, fmt.Errorf("GetUserURL: failed to scan row: %w", err)
		}

		userURL = append(userURL, model.GetUserURLResponse{
			ShortURL:    shortUrl,
			OriginalURL: url,
		})
	}
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("GetUserURL: something went wrong: %w", err)
	}

	return userURL, nil
}

func (s *Storage) SaveURL(ctx context.Context, userID int32, shortUrl string, originalUrl string) error {
	stmt, err := s.db.PrepareContext(ctx, `INSERT INTO url (short_url, url, user_id) VALUES ($1, $2, $3)`)
	if err != nil {
		return fmt.Errorf("SaveURL: failed to prepare query: %w", err)
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, shortUrl, originalUrl, userID)
	if err != nil {
		var pgxErr *pgconn.PgError
		if errors.As(err, &pgxErr) {
			if pgxErr.Code == pgerrcode.UniqueViolation {
				if pgxErr.ConstraintName == originalUrlConstraintName {
					// коллизия по url
					return errs.ErrOriginalUrlAlreadyExists
				} else if pgxErr.ConstraintName == shortUrlConstraintName {
					// коллизия по short url
					return errs.ErrShortUrlAlreadyExists
				}
			}
		}
		return fmt.Errorf("SaveURL: failed to save url: %w", err)
	}
	return nil
}

func (s *Storage) SaveURLBatch(ctx context.Context, userId int32, urls []model.URL) error {

	tx, err := s.db.BeginTx(ctx, nil)
	if err != nil {
		return fmt.Errorf("SaveURLBatch: failed to begin transaction: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.PrepareContext(ctx,
		`INSERT INTO url (short_url, url, user_id) VALUES ($1, $2, $3) 
			ON CONFLICT ON CONSTRAINT url_unique DO UPDATE SET url = EXCLUDED.url 
			RETURNING short_url, (xmax != 0)`)
	if err != nil {
		return fmt.Errorf("SaveURLBatch: failed to prepare query: %w", err)
	}
	defer stmt.Close()

	for i, url := range urls {

		var shortUrl string
		var isExist bool
		err = stmt.QueryRowContext(ctx, url.ShortURL, url.OriginalURL, userId).Scan(&shortUrl, &isExist)

		if err != nil {
			var pgxErr *pgconn.PgError
			if errors.As(err, &pgxErr) {
				if pgxErr.Code == pgerrcode.UniqueViolation {
					if pgxErr.ConstraintName == shortUrlConstraintName {
						// коллизия по short url
						return errs.ErrShortUrlAlreadyExists
					}
				}
			}
			return fmt.Errorf("SaveURLBatch: failed to save url: %w", err)
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
		return fmt.Errorf("SaveURLBatch: failed to commit transaction: %w", err)
	}

	return nil
}

func (s *Storage) Close() error {
	return s.db.Close()
}

func (s *Storage) Ping(ctx context.Context) error {
	err := s.db.PingContext(ctx)

	if err != nil {
		return fmt.Errorf("Ping: failed to ping db: %w ", err)
	}

	return nil
}
