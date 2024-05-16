package model

import (
	"fmt"
	"net/http"
	"regexp"
)

type CEP string

const invalidCEP CEP = "invalid CEP"

func BuildCEP(stringCep string) (CEP, *CustomError) {
	cep := CEP(stringCep)
	if cep.isValidCEP() {
		return cep, nil
	}

	return invalidCEP, NewCustomError(http.StatusBadRequest, fmt.Sprintf("CEP [%s] is not valid", stringCep))
}

func (cep CEP) isValidCEP() bool {
	regex := regexp.MustCompile(`^\d{5}-?\d{3}$`)
	return regex.MatchString(cep.ToString())
}

func (cep CEP) ToString() string {
	return string(cep)
}
