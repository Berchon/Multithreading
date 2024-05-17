package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/Berchon/Multithreading/internal/business/gateway"
	"github.com/Berchon/Multithreading/internal/business/model"
	"github.com/Berchon/Multithreading/internal/infra/service/config"
)

type GetAddressByCepUsecase interface {
	GetAddressByCep(cep model.CEP) (*model.AddressResponse, *model.CustomError)
}

type getAddressByCepUsecase struct {
	AddressService gateway.AddressService
}

func NewGetAddressByCepUsecase(addressService gateway.AddressService) GetAddressByCepUsecase {
	return &getAddressByCepUsecase{
		AddressService: addressService,
	}
}

func (getAddressByCepUsecase *getAddressByCepUsecase) GetAddressByCep(cep model.CEP) (*model.AddressResponse, *model.CustomError) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second)
	defer cancel()

	endpoint1 := config.NewEndpoint()
	endpoint2 := config.NewEndpoint()
	endpoint1.SetUrl("http://viacep.com.br/ws/%s/json/", cep)
	endpoint2.SetUrl("https://brasilapi.com.br/api/cep/v1/%s", cep)

	ch1 := make(chan []byte)
	ch2 := make(chan []byte)

	go getAddressByCepUsecase.AddressService.GetAddressByCep(ctx, endpoint1, ch1)
	go getAddressByCepUsecase.AddressService.GetAddressByCep(ctx, endpoint2, ch2)

	var address []byte
	var addressResponse *model.AddressResponse

	select {
	case address = <-ch1:
		var viaCep model.ViaCEP
		err := json.Unmarshal(address, &viaCep)
		if err != nil {
			log.Printf("Cep not found: %v", err)
			return nil, model.NewCustomError(http.StatusNotFound, fmt.Sprintf("error unmarshaling json: %v: CEP Not Found", err))
		}
		addressResponse = buildAddressFromViaCep(viaCep, "ViaCep")
	case address = <-ch2:
		var brasilAPI model.BrasilAPI
		if err := json.Unmarshal(address, &brasilAPI); err != nil {
			log.Printf("Cep not found: %v", err)
			return nil, model.NewCustomError(http.StatusNotFound, fmt.Sprintf("error unmarshaling json: %v: CEP Not Found", err))
		}
		addressResponse = buildAddressFromBrasilAPI(brasilAPI, "BrasilAPI")
	case <-time.After(1 * time.Second):
		log.Printf("Timeout to wait for the response from the APIs")
		return nil, model.NewCustomError(http.StatusGatewayTimeout, "Timeout to wait for the response from the APIs")
	}

	log.Printf("%s => %s-%s CEP: %s %s/%s\n", addressResponse.ApiName, addressResponse.Street, addressResponse.Neighborhood, addressResponse.ZipCode, addressResponse.City, addressResponse.State)
	return addressResponse, nil
}

func buildAddressFromViaCep(address model.ViaCEP, apiName string) *model.AddressResponse {
	return &model.AddressResponse{
		ApiName:      apiName,
		ZipCode:      address.ZipCode,
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
	}
}

func buildAddressFromBrasilAPI(address model.BrasilAPI, apiName string) *model.AddressResponse {
	return &model.AddressResponse{
		ApiName:      apiName,
		ZipCode:      address.ZipCode,
		Street:       address.Street,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
	}
}
