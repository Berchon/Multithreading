package validate

import (
	"net/http"

	"github.com/Berchon/Multithreading/internal/business/model"
	"github.com/go-chi/chi/v5"
)

const (
	cepPathParam = "cep"
)

func Cep(r *http.Request) (model.CEP, *model.CustomError) {
	cepParam := chi.URLParam(r, cepPathParam)
	if len(cepParam) == 0 {
		return "", model.NewCustomError(http.StatusBadRequest, "CEP path param is not precent")
	}
	cep, err := model.BuildCEP(cepParam)
	return cep, err
}
