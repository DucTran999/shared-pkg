package cache

import "errors"

var (
	ErrKeyNotFound = errors.New("key not found in cache")
	ErrMissingHost = errors.New("missing host")
)
