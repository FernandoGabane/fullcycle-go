package middleware

import (
    "context"
    "go-rate-limiter/internal/limiter"
    "net/http"
)

func RateLimitMiddleware(l *limiter.RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            ctx := context.Background()
            token := r.Header.Get("API_KEY")
            ip := r.RemoteAddr

            if token != "" {
                allowed, _ := l.Allow(ctx, "token:"+token, l.Cfg.TokenMaxReq)
                if !allowed {
                    http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
                    return
                }
            } else {
                allowed, _ := l.Allow(ctx, "ip:"+ip, l.Cfg.IPMaxReq)
                if !allowed {
                    http.Error(w, "you have reached the maximum number of requests or actions allowed within a certain time frame", http.StatusTooManyRequests)
                    return
                }
            }
            next.ServeHTTP(w, r)
        })
    }
}
