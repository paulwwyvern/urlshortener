package errs

import "errors"

var ErrShortUrlNotFound = errors.New("short url not found")
var ErrShortUrlAlreadyExists = errors.New("short url already exists")
