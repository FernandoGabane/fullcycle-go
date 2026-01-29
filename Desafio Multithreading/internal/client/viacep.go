package client

import (
	"context"
	"encoding/json"
	"fullcycle/go/exercise-2/internal/model"
	"net/http"
)

type ViaCepClient struct{}

func (v *ViaCepClient) Name() string {
	return "ViaCEP"
}

func (v *ViaCepClient) Fetch(ctx context.Context, cep string) (*model.Address, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"http://viacep.com.br/ws/"+cep+"/json/", nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		CEP        string `json:"cep"`
		Logradouro string `json:"logradouro"`
		Localidade string `json:"localidade"`
		UF         string `json:"uf"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &model.Address{
		CEP:    data.CEP,
		Street: data.Logradouro,
		City:   data.Localidade,
		State:  data.UF,
		Source: v.Name(),
	}, nil
}
