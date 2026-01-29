package main

import (
	"context"
	"net/http"
	"service-a/internal/handler"
	"service-a/internal/observability"
)

func main() {
	shutdown := observability.Init("service-a")
	defer shutdown(context.Background())
	http.HandleFunc("/cep", handler.Handle)
	http.ListenAndServe(":8080", nil)
}
