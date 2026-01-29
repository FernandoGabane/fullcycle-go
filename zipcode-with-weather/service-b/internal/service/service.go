package service

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"os"
	"service-b/internal/domain"

	"go.opentelemetry.io/otel"
)

func GetWeather(ctx context.Context, cep string) (*domain.WeatherResponse, error) {
	tracer := otel.Tracer("service-b")
	ctx, span := tracer.Start(ctx, "get-city")
	defer span.End()
	city, err := getCity(cep)
	if err != nil {
		return nil, err
	}
	ctx, span2 := tracer.Start(ctx, "get-temp")
	defer span2.End()
	temp, err := getTemp(city)
	if err != nil {
		return nil, err
	}
	return &domain.WeatherResponse{
		City:  city,
		TempC: temp,
		TempF: temp*1.8 + 32,
		TempK: temp + 273,
	}, nil
}

func getCity(cep string) (string, error) {
	r, _ := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	defer r.Body.Close()
	var v struct {
		Localidade string `json:"localidade"`
	}
	json.NewDecoder(r.Body).Decode(&v)
	if v.Localidade == "" {
		return "", errors.New("not found")
	}
	return v.Localidade, nil
}

func getTemp(city string) (float64, error) {
	r, _ := http.Get("http://api.weatherapi.com/v1/current.json?key=" + os.Getenv("WEATHER_API_KEY") + "&q=" + city)
	defer r.Body.Close()
	var w struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	json.NewDecoder(r.Body).Decode(&w)
	return w.Current.TempC, nil
}
