
package service

import (
    "encoding/json"
    "errors"
    "net/http"
)

type ViaCEPService struct{}

func NewViaCEPService() *ViaCEPService {
    return &ViaCEPService{}
}

func (v *ViaCEPService) ResolveCity(cep string) (string, error) {
    resp, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
    if err != nil {
        return "", err
    }
    defer resp.Body.Close()

    var data struct {
        Localidade string `json:"localidade"`
        Erro bool `json:"erro"`
    }
    json.NewDecoder(resp.Body).Decode(&data)

    if data.Erro || data.Localidade == "" {
        return "", errors.New("not found")
    }
    return data.Localidade, nil
}
