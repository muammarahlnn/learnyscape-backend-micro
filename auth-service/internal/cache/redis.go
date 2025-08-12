package cache

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type Redis interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key string, value any, ttl time.Duration) error
	Delete(ctx context.Context, keys ...string) error
	Pipeline(ctx context.Context, fn func(ctx context.Context, pipe redis.Pipeliner) error) error
	Exists(ctx context.Context, keys ...string) (int64, error)
	Expire(ctx context.Context, key string, ttl time.Duration) (bool, error)
}

type redisClusterImpl struct {
	rdb *redis.ClusterClient
}

func NewRedisCluster(rdb *redis.ClusterClient) Redis {
	return &redisClusterImpl{
		rdb: rdb,
	}
}

func (r *redisClusterImpl) Get(ctx context.Context, key string) (string, error) {
	return r.rdb.Get(ctx, key).Result()
}

func (r *redisClusterImpl) Set(ctx context.Context, key string, value any, ttl time.Duration) error {
	return r.rdb.Set(ctx, key, value, ttl).Err()
}

func (r *redisClusterImpl) Delete(ctx context.Context, keys ...string) error {
	return r.rdb.Del(ctx, keys...).Err()
}

func (r *redisClusterImpl) Pipeline(ctx context.Context, fn func(ctx context.Context, pipe redis.Pipeliner) error) error {
	pipe := r.rdb.Pipeline()
	if err := fn(ctx, pipe); err != nil {
		return err
	}

	_, err := pipe.Exec(ctx)
	return err
}

func (r *redisClusterImpl) Exists(ctx context.Context, keys ...string) (int64, error) {
	return r.rdb.Exists(ctx, keys...).Result()
}

func (r *redisClusterImpl) Expire(ctx context.Context, key string, ttl time.Duration) (bool, error) {
	return r.rdb.Expire(ctx, key, ttl).Result()
}
