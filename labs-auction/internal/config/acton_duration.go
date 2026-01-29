package config

import (
	"os"
	"strconv"
	"time"
)

func GetAuctionDuration() time.Duration {
	value := os.Getenv("AUCTION_DURATION_SECONDS")
	seconds, err := strconv.Atoi(value)
	if err != nil || seconds <= 0 {
		return 60 * time.Second // fallback padrão
	}
	return time.Duration(seconds) * time.Second
}
