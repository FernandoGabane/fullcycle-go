package main

import (
    "go-rate-limiter/internal/config"
    "go-rate-limiter/internal/limiter"
    "go-rate-limiter/internal/middleware"
    "go-rate-limiter/internal/storage/redisstore"
    "log"
    "net/http"
)

func main() {
    cfg := config.Load()
    store := redisstore.NewRedisStore(cfg.RedisAddr)
    limiterService := limiter.NewRateLimiter(store, cfg)

    mux := http.NewServeMux()
    mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        w.Write([]byte("OK"))
    })

    handler := middleware.RateLimitMiddleware(limiterService)(mux)
    log.Println("Server running on :8080")
    http.ListenAndServe(":8080", handler)
}
