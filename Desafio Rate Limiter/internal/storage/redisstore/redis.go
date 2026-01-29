package redisstore

import (
    "context"
    "time"
    "github.com/go-redis/redis/v8"
)

type RedisStore struct {
    client *redis.Client
}

func NewRedisStore(addr string) *RedisStore {
    return &RedisStore{
        client: redis.NewClient(&redis.Options{Addr: addr}),
    }
}

func (r *RedisStore) Increment(ctx context.Context, key string, window time.Duration, block time.Duration) (int, error) {
    pipe := r.client.TxPipeline()
    incr := pipe.Incr(ctx, key)
    pipe.Expire(ctx, key, block)
    _, err := pipe.Exec(ctx)
    if err != nil {
        return 0, err
    }
    return int(incr.Val()), nil
}
