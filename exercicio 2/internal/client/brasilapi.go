package client

import (
	"context"
	"encoding/json"
	"fullcycle/go/exercise-2/internal/model"
	"net/http"
)

type BrasilAPIClient struct{}

func (b *BrasilAPIClient) Name() string {
	return "BrasilAPI"
}

func (b *BrasilAPIClient) Fetch(ctx context.Context, cep string) (*model.Address, error) {
	req, err := http.NewRequestWithContext(ctx, http.MethodGet,
		"https://brasilapi.com.br/api/cep/v1/"+cep, nil)
	if err != nil {
		return nil, err
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var data struct {
		CEP    string `json:"cep"`
		Street string `json:"street"`
		City   string `json:"city"`
		State  string `json:"state"`
	}

	if err := json.NewDecoder(resp.Body).Decode(&data); err != nil {
		return nil, err
	}

	return &model.Address{
		CEP:    data.CEP,
		Street: data.Street,
		City:   data.City,
		State:  data.State,
		Source: b.Name(),
	}, nil
}
