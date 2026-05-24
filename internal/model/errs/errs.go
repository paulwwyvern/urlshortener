package errs

import "errors"

var ErrShortUrlNotFound = errors.New("short url not found")
var ErrShortUrlGone = errors.New("short url gone")
var ErrOriginalUrlNotFound = errors.New("original url not found")

var ErrShortUrlAlreadyExists = errors.New("short url already exists")
var ErrOriginalUrlAlreadyExists = errors.New("original url already exists")

//var ErrInternalError = errors.New("internal error")

var ErrShortUrlForbidden = errors.New("short url forbidden")
