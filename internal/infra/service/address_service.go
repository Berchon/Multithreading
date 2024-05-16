package service

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"github.com/Berchon/Multithreading/internal/business/gateway"
	"github.com/Berchon/Multithreading/internal/business/model"
	"github.com/Berchon/Multithreading/internal/infra/service/config"
)

type addressService struct {
	endpoint *config.Endpoint
}

func NewAddressService() gateway.AddressService {
	endpoint := config.NewEndpoint()
	return &addressService{
		endpoint: endpoint,
	}
}

func (addressService *addressService) GetAddressByCep(ctx context.Context, endpoint *config.Endpoint, ch chan []byte) *model.CustomError {
	url := endpoint.GetUrl()
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)
	if err != nil {
		ch <- nil
		return model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error creating request: %v", err))
	}

	response, err := http.DefaultClient.Do(req)
	if err != nil {
		ch <- nil
		return model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error sending request: %v", err))
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		ch <- nil
		return model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error reading response: %v", err))
	}

	if response.StatusCode != http.StatusOK {
		ch <- nil
		return model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error reading response. Status code not OK: Status code returned %v", response.StatusCode))
	}

	if string(body) == "{\n  \"erro\": true\n}" {
		ch <- nil
		return model.NewCustomError(http.StatusInternalServerError, fmt.Sprintf("error reading response. Status code not OK: Status code returned %v", response.StatusCode))
	}

	ch <- body
	return nil
}
