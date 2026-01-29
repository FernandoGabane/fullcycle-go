package limiter

import (
    "context"
    "go-rate-limiter/internal/config"
    "go-rate-limiter/internal/storage"
    "time"
)

type RateLimiter struct {
    Store storage.Store
    Cfg   *config.Config
}

func NewRateLimiter(store storage.Store, cfg *config.Config) *RateLimiter {
    return &RateLimiter{Store: store, Cfg: cfg}
}

func (r *RateLimiter) Allow(ctx context.Context, key string, max int) (bool, error) {
    count, err := r.Store.Increment(ctx, key, time.Second, time.Duration(r.Cfg.BlockDuration)*time.Second)
    if err != nil {
        return false, err
    }
    return count <= max, nil
}
