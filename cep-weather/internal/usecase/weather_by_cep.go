
package usecase

import (
    "encoding/json"
    "net/http"
    "regexp"
    "cepweather/internal/service"
)

type WeatherByCEP struct {
    cep service.CEPService
    weather service.WeatherService
}

func NewWeatherByCEP() *WeatherByCEP {
    return &WeatherByCEP{
        cep: service.NewViaCEPService(),
        weather: service.NewWeatherAPIService(),
    }
}

func (u *WeatherByCEP) Handle(w http.ResponseWriter, r *http.Request) {
    cep := r.URL.Query().Get("cep")
    if !regexp.MustCompile(`^\d{8}$`).MatchString(cep) {
        http.Error(w, "invalid zipcode", http.StatusUnprocessableEntity)
        return
    }

    city, err := u.cep.ResolveCity(cep)
    if err != nil {
        http.Error(w, "can not find zipcode", http.StatusNotFound)
        return
    }

    tempC, err := u.weather.GetTemperature(city)
    if err != nil {
        http.Error(w, "weather error", http.StatusInternalServerError)
        return
    }

    resp := map[string]float64{
        "temp_C": tempC,
        "temp_F": tempC*1.8 + 32,
        "temp_K": tempC + 273,
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(resp)
}
