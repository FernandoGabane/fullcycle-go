package main

import (
	"context"
	"net/http"
	"service-b/internal/handler"
	"service-b/internal/observability"
)

func main() {
	shutdown := observability.Init("service-b")
	defer shutdown(context.Background())
	http.HandleFunc("/weather", handler.Handle)
	http.ListenAndServe(":8081", nil)
}
