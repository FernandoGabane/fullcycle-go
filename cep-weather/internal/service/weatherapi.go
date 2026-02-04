package service

import (
	"encoding/json"
	"errors"
	"net/http"
	"net/url"
	"os"
)

type WeatherAPIService struct{}

func NewWeatherAPIService() *WeatherAPIService {
	return &WeatherAPIService{}
}

func (w *WeatherAPIService) GetTemperature(city string) (float64, error) {
	key := os.Getenv("WEATHER_API_KEY")

	encodedCity := url.QueryEscape(city)

	endpoint := "https://api.weatherapi.com/v1/current.json?key=" +
		key + "&q=" + encodedCity

	resp, err := http.Get(endpoint)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var data struct {
		Current struct {
			TempC float64 `json:"temp_c"`
		} `json:"current"`
	}
	json.NewDecoder(resp.Body).Decode(&data)

	if data.Current.TempC == 0 {
		return 0, errors.New("not found")
	}
	return data.Current.TempC, nil
}
