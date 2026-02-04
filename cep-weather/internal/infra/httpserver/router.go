
package httpserver

import (
    "net/http"
    "cepweather/internal/usecase"
)

func NewRouter() http.Handler {
    mux := http.NewServeMux()
    uc := usecase.NewWeatherByCEP()
    mux.HandleFunc("/weather", uc.Handle)
    return mux
}
