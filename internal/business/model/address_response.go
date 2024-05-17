package model

type AddressResponse struct {
	ApiName      string `json:"api_name"`
	ZipCode      string `json:"zip_code"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
}
