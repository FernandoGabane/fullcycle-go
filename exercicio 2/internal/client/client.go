package client

import (
	"context"
	"fullcycle/go/exercise-2/internal/model"
)

type CepClient interface {
	Fetch(ctx context.Context, cep string) (*model.Address, error)
	Name() string
}
