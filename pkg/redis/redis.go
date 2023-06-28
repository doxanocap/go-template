package redis

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

type Config struct {
	Host      string
	Password  string
	DB        int
	KeyPrefix string
}

type Conn struct {
	client    *redis.Client
	keyPrefix string
}

func Connect(ctx context.Context, config Config) (*Conn, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     config.Host,
		Password: config.Password,
		DB:       config.DB,
	})

	if err := client.Ping(ctx).Err(); err != nil {
		return nil, fmt.Errorf("failed to connect: %v", err)
	}

	return &Conn{
		client:    client,
		keyPrefix: config.KeyPrefix,
	}, nil
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
