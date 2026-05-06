package errs

import "errors"

var ErrShortUrlNotFound = errors.New("short url not found")
var ErrShortUrlAlreadyExists = errors.New("short url already exists")
var ErrOriginalUrlAlreadyExists = errors.New("original url already exists")
var ErrInternalError = errors.New("internal error")
var ErrStorageUnavailable = errors.New("storage unavailable")
