package handler

import (
	"net/http"

	"github.com/Berchon/Multithreading/internal/business/usecase"
	"github.com/Berchon/Multithreading/internal/infra/webapp/request/validate"
)

type GetAddressByCepHandler interface {
	HttpHandler
}

type getAddressByCepHandler struct {
	usecase  usecase.GetAddressByCepUsecase
	response *responseHandler
}

func NewGetAddressByCepHandler(usecase usecase.GetAddressByCepUsecase) GetAddressByCepHandler {
	response := NewResponseHandler()
	return &getAddressByCepHandler{
		usecase:  usecase,
		response: response,
	}
}

// Get Address by CEP
// @Summary Get Address by CEP
// @Description Get Address information by CEP
// @Tags CEP
// @Accept json
// @Produce json
// @Param cep path string true "CEP to be queried. Must be composed of 8 numeric digits." Format(CEP)
// @Success 200 {object} model.AddressResponse "Address information"
// @Failure 400 {object} model.CustomError "Validation CEP error"
// @Failure 404 {object} model.CustomError "CEP not found"
// @Failure 504 {object} model.CustomError "Timeout error"
// @Router /{cep} [get]
func (getAddressByCepHandler *getAddressByCepHandler) Handle(w http.ResponseWriter, r *http.Request) {
	cep, err := validate.Cep(r)
	if err != nil {
		getAddressByCepHandler.response.RequestResponse(w, r, err, err.StatusCode)
		return
	}

	address, err := getAddressByCepHandler.usecase.GetAddressByCep(cep)
	if err != nil {
		getAddressByCepHandler.response.RequestResponse(w, r, err, err.StatusCode)
		return
	}

	getAddressByCepHandler.response.RequestResponse(w, r, address, http.StatusOK)
}
