package main

import (
	"flag"
	"fmt"
	"go-load-tester/internal/runner"
	"time"
)

func main() {
	url := flag.String("url", "", "Target URL")
	requests := flag.Int("requests", 1, "Total number of requests")
	concurrency := flag.Int("concurrency", 1, "Concurrent requests")

	flag.Parse()

	if *url == "" {
		panic("url is required")
	}

	start := time.Now()
	report := runner.Run(*url, *requests, *concurrency)
	elapsed := time.Since(start)

	fmt.Println("==== Load Test Report ====")
	fmt.Println("Total time:", elapsed)
	fmt.Println("Total requests:", report.Total)
	fmt.Println("Status 200:", report.Success)
	fmt.Println("Other status codes:")
	for code, count := range report.Errors {
		fmt.Printf("  %d: %d", code, count)
	}
}
