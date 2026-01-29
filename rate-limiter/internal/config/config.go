package config

import (
    "fmt"
    "os"
    "github.com/joho/godotenv"
)

type Config struct {
    RedisAddr     string
    IPMaxReq      int
    TokenMaxReq   int
    BlockDuration int
}

func Load() *Config {
    _ = godotenv.Load()
    return &Config{
        RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
        IPMaxReq:      getEnvInt("IP_MAX_REQ", 5),
        TokenMaxReq:   getEnvInt("TOKEN_MAX_REQ", 10),
        BlockDuration: getEnvInt("BLOCK_DURATION_SECONDS", 300),
    }
}

func getEnv(key, def string) string {
    if v := os.Getenv(key); v != "" {
        return v
    }
    return def
}

func getEnvInt(key string, def int) int {
    if v := os.Getenv(key); v != "" {
        var i int
        fmt.Sscanf(v, "%d", &i)
        return i
    }
    return def
}
