package cache

import (
	"context"
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/redis/go-redis/v9"
)

type redisCache struct {
	client *redis.Client
}

func NewRedisCache(config Config) (Cache, error) {
	// Initialize Redis client
	rdb := redis.NewClient(&redis.Options{
		Addr:     net.JoinHostPort(config.Host, strconv.Itoa(config.Port)),
		Password: config.Password,
		DB:       config.DB,

		// Set connection pool options
		DialTimeout:  5 * time.Second,
		ReadTimeout:  3 * time.Second,
		WriteTimeout: 3 * time.Second,
		PoolSize:     10,
		MinIdleConns: 5,
	})

	// Test the connection
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := rdb.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect to redis: %w", err)
	}

	return &redisCache{
		client: rdb,
	}, nil
}

func (r *redisCache) Get(ctx context.Context, key string) (string, error) {
	val, err := r.client.Get(ctx, key).Result()
	if err != nil {
		// Match the error pattern used in the in-memory implementation for consistency
		if err == redis.Nil {
			return "", fmt.Errorf("%w: %s", ErrKeyNotFound, key)
		}

		return "", err
	}

	return val, nil
}

func (r *redisCache) Set(ctx context.Context, key string, value any, expiration time.Duration) error {
	strVal := fmt.Sprintf("%v", value)
	if strVal == "" {
		return fmt.Errorf("value cannot be empty")
	}

	return r.client.Set(ctx, key, strVal, expiration).Err()
}

func (r *redisCache) Del(ctx context.Context, keys ...string) error {
	return r.client.Del(ctx, keys...).Err()
}

func (r *redisCache) Ping(ctx context.Context) error {
	return r.client.Ping(ctx).Err()
}

func (r *redisCache) Close() error {
	return r.client.Close()
}
