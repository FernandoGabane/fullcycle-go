package service

import (
	"context"
	"errors"
	"fullcycle/go/exercise-2/internal/client"
	"fullcycle/go/exercise-2/internal/model"
	"time"
)

type CepService struct {
	clients []client.CepClient
}

func NewCepService() *CepService {
	return &CepService{
		clients: []client.CepClient{
			&client.BrasilAPIClient{},
			&client.ViaCepClient{},
		},
	}
}

func (s *CepService) GetFastest(cep string) (*model.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	ch := make(chan *model.Address, len(s.clients))

	for _, c := range s.clients {
		go func(cl client.CepClient) {
			if res, err := cl.Fetch(ctx, cep); err == nil {
				ch <- res
			}
		}(c)
	}

	select {
	case res := <-ch:
		return res, nil
	case <-ctx.Done():
		return nil, errors.New("timeout: no API responded within 1 second")
	}
}
