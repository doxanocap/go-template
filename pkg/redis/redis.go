package redis

import (
	"app/internal/model"
	"context"
	"github.com/doxanocap/pkg/lg"
	"github.com/go-redis/redis/v8"
	"time"
)

type Conn struct {
	client    *redis.Client
	keyPrefix string
}

func InitConnection(cfg *model.Config) *Conn {
	ctx := context.Background()
	client := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Host,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.Database,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		lg.Fatalf("redis connection: %v", err)
	}

	return &Conn{
		client:    client,
		keyPrefix: cfg.Redis.Prefix,
	}
}

func (r *Conn) Get(ctx context.Context, key string) ([]byte, error) {
	key = r.keyPrefix + key
	return r.client.Get(ctx, key).Bytes()
}

func (r *Conn) Set(ctx context.Context, key string, value []byte) error {
	key = r.keyPrefix + key
	return r.client.Set(ctx, key, value, 0).Err()
}

func (r *Conn) SetWithTTL(ctx context.Context, key string, value []byte, ttl time.Duration) error {
	key = r.keyPrefix + key
	return r.client.Set(ctx, key, value, ttl).Err()
}

func (r *Conn) Delete(ctx context.Context, key string) error {
	key = r.keyPrefix + key
	return r.client.Del(ctx, key).Err()
}

func (r *Conn) FlushAll(ctx context.Context) error {
	return r.client.FlushAll(ctx).Err()
}

func (r *Conn) Close() error {
	return r.client.Close()
}
