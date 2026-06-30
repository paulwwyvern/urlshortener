package errs

import (
	"errors"
	"strings"
)

var ErrShortUrlNotFound = errors.New("short url not found")
var ErrShortUrlGone = errors.New("short url gone")
var ErrOriginalUrlNotFound = errors.New("original url not found")

var ErrShortUrlAlreadyExists = errors.New("short url already exists")
var ErrOriginalUrlAlreadyExists = errors.New("original url already exists")

//var ErrInternalError = errors.New("internal error")

var ErrShortUrlForbidden = errors.New("short url forbidden")

type ErrAuditNotify struct {
	errs []error
}

func NewErrAuditNotify(errs ...error) *ErrAuditNotify {
	return &ErrAuditNotify{errs: errs}
}

func (e *ErrAuditNotify) Error() string {
	errsString := make([]string, 0, len(e.errs))

	for _, err := range e.errs {
		errsString = append(errsString, err.Error())
	}

	return strings.Join(errsString, "; ")
}

func (e *ErrAuditNotify) Unwrap() []error {
	return e.errs
}
