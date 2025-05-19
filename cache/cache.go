package cache

import (
	"context"
	"time"
)

type Cache interface {
	Get(ctx context.Context, key string) (string, error)

	Set(ctx context.Context, key string, value any, expiration time.Duration) error

	Del(ctx context.Context, keys ...string) error

	Ping(ctx context.Context) error

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
	if config.IsCacheOnMemory {
		return NewRistrettoCache()
	}

	return NewRedisCache(config)
}
