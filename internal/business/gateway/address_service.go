package gateway

import (
	"context"

	"github.com/Berchon/Multithreading/internal/business/model"
	"github.com/Berchon/Multithreading/internal/infra/service/config"
)

type AddressService interface {
	GetAddressByCep(ctx context.Context, endpoint *config.Endpoint, ch chan []byte) *model.CustomError
}
