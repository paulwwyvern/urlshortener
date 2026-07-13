package postgres

import (
	"context"
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/paulwwyvern/urlshortener/internal/model"
	"github.com/paulwwyvern/urlshortener/internal/model/dto"
	"github.com/paulwwyvern/urlshortener/internal/model/errs"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestStorage_GetURL(t *testing.T) {
	url := "http://example.com"
	shortURL := "abcdef"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"url", "is_deleted"}).AddRow(url, false)

	query := regexp.QuoteMeta("SELECT url, is_deleted FROM url")
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(shortURL).WillReturnRows(rows)

	storage := &Storage{db: db}
	res, err := storage.GetURL(context.Background(), shortURL)
	assert.NoError(t, err)
	assert.Equal(t, url, res)
}

func TestStorage_GetURL_NotFound(t *testing.T) {
	shortURL := "abcdef"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"url", "is_deleted"})

	query := regexp.QuoteMeta("SELECT url, is_deleted FROM url")
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(shortURL).WillReturnRows(rows)

	storage := &Storage{db: db}
	_, err = storage.GetURL(context.Background(), shortURL)
	assert.ErrorIs(t, err, errs.ErrShortURLGone)
}

func TestStorage_GetURL_Gone(t *testing.T) {
	shortURL := "abcdef"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"url", "is_deleted"}).AddRow("", true)

	query := regexp.QuoteMeta("SELECT url, is_deleted FROM url")
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(shortURL).WillReturnRows(rows)

	storage := &Storage{db: db}
	_, err = storage.GetURL(context.Background(), shortURL)
	assert.ErrorIs(t, err, errs.ErrShortURLGone)
}

func TestStorage_GetShortURL(t *testing.T) {
	url := "http://example.com"
	shortURL := "abcdef"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"short_url", "is_deleted"}).AddRow(shortURL, false)

	query := regexp.QuoteMeta("SELECT short_url, is_deleted FROM url")
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(url).WillReturnRows(rows)

	storage := &Storage{db: db}
	res, err := storage.GetShortURL(context.Background(), url)
	assert.NoError(t, err)
	assert.Equal(t, shortURL, res)
}

func TestStorage_GetShortURL_NotFound(t *testing.T) {

	url := "http://example.com"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"short_url", "is_deleted"})

	query := regexp.QuoteMeta("SELECT short_url, is_deleted FROM url")
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(url).WillReturnRows(rows)

	storage := &Storage{db: db}
	_, err = storage.GetShortURL(context.Background(), url)
	assert.ErrorIs(t, err, errs.ErrShortURLNotFound)
}

func TestStorage_GetShortURL_Gone(t *testing.T) {

	url := "http://example.com"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"short_url", "is_deleted"}).AddRow("", true)

	query := regexp.QuoteMeta("SELECT short_url, is_deleted FROM url")
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(url).WillReturnRows(rows)

	storage := &Storage{db: db}
	_, err = storage.GetShortURL(context.Background(), url)
	assert.ErrorIs(t, err, errs.ErrShortURLNotFound)
}

func TestStorage_GetUserURL(t *testing.T) {
	userID := int32(123456)

	urls := []dto.GetUserURLResponse{
		{
			ShortURL:    "a",
			OriginalURL: "http://a.com",
		},
		{
			ShortURL:    "b",
			OriginalURL: "http://b.com",
		},
		{
			ShortURL:    "d",
			OriginalURL: "http://d.com",
		},
	}

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rows := sqlmock.NewRows([]string{"short_url", "url", "is_deleted"}).
		AddRow("a", "http://a.com", false).
		AddRow("b", "http://b.com", false).
		AddRow("c", "http://c.com", true).
		AddRow("d", "http://d.com", false)

	query := regexp.QuoteMeta("SELECT short_url, url, is_deleted FROM url")
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs(userID).WillReturnRows(rows)

	storage := &Storage{db: db}
	res, err := storage.GetUserURL(context.Background(), userID)
	assert.NoError(t, err)
	assert.Equal(t, urls, res)
}

func TestStorage_SaveURL(t *testing.T) {
	userID := int32(123456)

	url := "http://example.com"
	shortURL := "abcdef"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	query := regexp.QuoteMeta("INSERT INTO url (short_url, url, user_id)")
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(shortURL, url, userID).WillReturnResult(sqlmock.NewResult(1, 1))

	storage := &Storage{db: db}
	err = storage.SaveURL(context.Background(), userID, shortURL, url)
	assert.NoError(t, err)
}

func TestStorage_SaveURL_OriginalURLConflict(t *testing.T) {
	userID := int32(123456)

	url := "http://example.com"
	shortURL := "abcdef"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	retErr := &pgconn.PgError{
		Code:           pgerrcode.UniqueViolation,
		ConstraintName: originalURLConstraintName,
	}

	query := regexp.QuoteMeta("INSERT INTO url (short_url, url, user_id)")
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(shortURL, url, userID).WillReturnError(retErr)

	storage := &Storage{db: db}
	err = storage.SaveURL(context.Background(), userID, shortURL, url)
	assert.ErrorIs(t, err, errs.ErrOriginalURLAlreadyExists)
}

func TestStorage_SaveURL_ShortURLConflict(t *testing.T) {
	userID := int32(123456)

	url := "http://example.com"
	shortURL := "abcdef"

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	retErr := &pgconn.PgError{
		Code:           pgerrcode.UniqueViolation,
		ConstraintName: shortURLConstraintName,
	}

	query := regexp.QuoteMeta("INSERT INTO url (short_url, url, user_id)")
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs(shortURL, url, userID).WillReturnError(retErr)

	storage := &Storage{db: db}
	err = storage.SaveURL(context.Background(), userID, shortURL, url)
	assert.ErrorIs(t, err, errs.ErrShortURLAlreadyExists)
}

func TestStorage_SaveURLBatch(t *testing.T) {
	userID := int32(123456)
	urls := []model.URL{
		{
			ShortURL:    "a",
			OriginalURL: "http://a.com",
		},
		{
			ShortURL:    "b",
			OriginalURL: "http://b.com",
		},
		{
			ShortURL:    "c",
			OriginalURL: "http://c.com",
		},
		{
			ShortURL:    "d",
			OriginalURL: "http://d.com",
		},
	}

	expectedURLs := []model.URL{
		{
			ShortURL:    "a",
			OriginalURL: "http://a.com",
		},
		{
			ShortURL:    "b",
			OriginalURL: "http://b.com",
		},
		{
			ShortURL:    "C",
			OriginalURL: "http://c.com",
			IsExist:     true,
		},
		{
			ShortURL:    "d",
			OriginalURL: "http://d.com",
		},
	}

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	rowsA := sqlmock.NewRows([]string{"short_url", "is_exist"}).AddRow("", false)
	rowsB := sqlmock.NewRows([]string{"short_url", "is_exist"}).AddRow("", false)
	rowsC := sqlmock.NewRows([]string{"short_url", "is_exist"}).AddRow("C", true)
	rowsD := sqlmock.NewRows([]string{"short_url", "is_exist"}).AddRow("", false)

	query := regexp.QuoteMeta("INSERT INTO url (short_url, url, user_id)")
	mock.ExpectBegin()
	mock.ExpectPrepare(query)
	mock.ExpectQuery(query).WithArgs("a", "http://a.com", userID).WillReturnRows(rowsA)
	mock.ExpectQuery(query).WithArgs("b", "http://b.com", userID).WillReturnRows(rowsB)
	mock.ExpectQuery(query).WithArgs("c", "http://c.com", userID).WillReturnRows(rowsC)
	mock.ExpectQuery(query).WithArgs("d", "http://d.com", userID).WillReturnRows(rowsD)
	mock.ExpectCommit()

	storage := &Storage{db: db}
	err = storage.SaveURLBatch(context.Background(), userID, urls)
	assert.NoError(t, err)
	assert.Equal(t, expectedURLs, urls)
}

func TestStorage_SoftDeleteURLBatch(t *testing.T) {
	userID := int32(123456)
	shortURLs := []string{"a", "b", "c", "d"}

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	query := regexp.QuoteMeta("UPDATE url SET is_deleted = TRUE")
	mock.ExpectBegin()
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs("a", userID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WithArgs("b", userID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WithArgs("c", userID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WithArgs("d", userID).WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	storage := &Storage{db: db}
	err = storage.SoftDeleteURLBatch(context.Background(), userID, shortURLs)
	assert.NoError(t, err)
}

func TestStorage_PurgeURLBatch(t *testing.T) {
	shortURLs := []string{"a", "b", "c", "d"}

	db, mock, err := sqlmock.New()
	require.NoError(t, err)
	defer db.Close()

	query := regexp.QuoteMeta("DELETE FROM url")
	mock.ExpectBegin()
	mock.ExpectPrepare(query)
	mock.ExpectExec(query).WithArgs("a").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WithArgs("b").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WithArgs("c").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectExec(query).WithArgs("d").WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	storage := &Storage{db: db}
	err = storage.PurgeURLBatch(context.Background(), shortURLs)
	assert.NoError(t, err)
}
