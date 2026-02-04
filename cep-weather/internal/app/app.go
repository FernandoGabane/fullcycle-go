
package app

import (
    "log"
    "net/http"
    "cepweather/internal/infra/httpserver"
)

func Start() {
    handler := httpserver.NewRouter()
    log.Fatal(http.ListenAndServe(":8080", handler))
}
