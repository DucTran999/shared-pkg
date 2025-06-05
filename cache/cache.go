package cache

import (
	"context"
	"time"
)

type Cache interface {
	// Get retrieves a value from the cache by its key.
	Get(ctx context.Context, key string) (string, error)

	// Set stores a value in the cache with an optional expiration time.
	Set(ctx context.Context, key string, value any, expiration time.Duration) error

	// Del removes one or more keys from the cache.
	Del(ctx context.Context, keys ...string) error

	// Ping checks the connection to the cache server.
	Ping(ctx context.Context) error

	// Close closes the cache connection.
	Close() error
}

type Config struct {
	IsCacheOnMemory bool

	Host     string
	Port     int
	Password string
	DB       int
}

func NewCache(config Config) (Cache, error) {
	// Validate configuration
	if !config.IsCacheOnMemory {
		if config.Host == "" {
			return nil, ErrMissingHost
		}
	}

	if config.IsCacheOnMemory {
		return NewRistrettoCache()
	}

	return NewRedisCache(config)
}
