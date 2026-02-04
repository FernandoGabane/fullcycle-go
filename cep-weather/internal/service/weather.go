
package service

type WeatherService interface {
    GetTemperature(city string) (float64, error)
}
