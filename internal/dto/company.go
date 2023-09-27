package dto

import (
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type CompanyDto struct {
	ID       string          `json:"id"`
	Name     string          `json:"name"`
	Document string          `json:"document"`
	Address  *entity.Address `json:"address"`
	Status   string          `json:"status"`
}

type AddressDto struct {
	ID           string `json:"id"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
	CompanyId    string `json:"company_id"`
}

type AddressResponseDto struct {
	ID           string `json:"id"`
	Street       string `json:"street"`
	Number       string `json:"number"`
	Complement   string `json:"complement"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Country      string `json:"country"`
	ZipCode      string `json:"zip_code"`
}
