package cache

import (
	"context"
	"fmt"
	"time"

	"github.com/dgraph-io/ristretto/v2"
)

type ristrettoCache struct {
	cache *ristretto.Cache[string, string]
}

func NewRistrettoCache() (Cache, error) {
	c, err := ristretto.NewCache(&ristretto.Config[string, string]{
		NumCounters: 1e7,     // number of keys to track frequency of (10M).
		MaxCost:     1 << 30, // maximum cost of cache (1GB).
		BufferItems: 64,      // number of keys per Get buffer.

	})

	if err != nil {
		return nil, err
	}

	return &ristrettoCache{cache: c}, nil
}

func (r *ristrettoCache) Get(ctx context.Context, key string) (string, error) {
	val, found := r.cache.Get(key)
	if !found {
		return "", fmt.Errorf("key not found: %s", key)
	}

	return val, nil
}

func (r *ristrettoCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	strVal := fmt.Sprintf("%v", value)
	if strVal == "" {
		return fmt.Errorf("value cannot be empty")
	}

	// Set the value in the cache with a TTL
	ok := r.cache.SetWithTTL(key, strVal, 1, expiration)
	if !ok {
		return fmt.Errorf("failed to set key: %s", key)
	}

	return nil
}

func (r *ristrettoCache) Del(ctx context.Context, keys ...string) error {
	for _, key := range keys {
		r.cache.Del(key)
	}

	return nil
}

func (r *ristrettoCache) Ping(ctx context.Context) error {
	// Ristretto does not have a ping method, but we can check if the cache is initialized
	if r.cache == nil {
		return fmt.Errorf("cache is not initialized")
	}

	return nil
}

func (r *ristrettoCache) Close() error {
	r.cache.Close()
	return nil
}
