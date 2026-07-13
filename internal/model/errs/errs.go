package errs

import (
	"errors"
)

var ErrShortURLNotFound = errors.New("short url not found")
var ErrShortURLGone = errors.New("short url gone")
var ErrOriginalURLNotFound = errors.New("original url not found")

var ErrShortURLAlreadyExists = errors.New("short url already exists")
var ErrOriginalURLAlreadyExists = errors.New("original url already exists")

//var ErrInternalError = errors.New("internal error")

var ErrShortURLForbidden = errors.New("short url forbidden")
