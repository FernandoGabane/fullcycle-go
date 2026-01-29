package storage

import (
    "context"
    "time"
)

type Store interface {
    Increment(ctx context.Context, key string, window time.Duration, block time.Duration) (int, error)
}
