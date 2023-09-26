package entity

import (
	"github.com/dedicio/sisgares-accounts-service/pkg/utils"
)

type CompanyRepository interface {
	FindById(id string) (*Company, error)
	FindAll() ([]*Company, error)
	Create(company *Company) error
	Update(company *Company) error
	Delete(id string) error
}

type AddressRepository interface {
	FindById(id string) (*Address, error)
	Create(category *Address) error
	Update(category *Address) error
	Delete(id string) error
}

type Address struct {
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

func NewAddress(
	street string,
	number string,
	complement string,
	neighborhood string,
	city string,
	state string,
	country string,
	zipCode string,
	companyId string,
) *Address {
	id := utils.NewUUID()
	return &Address{
		ID:           id,
		Street:       street,
		Number:       number,
		Complement:   complement,
		Neighborhood: neighborhood,
		City:         city,
		State:        state,
		Country:      country,
		ZipCode:      zipCode,
		CompanyId:    companyId,
	}
}

type Company struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Document string   `json:"document"`
	Address  *Address `json:"address"`
	Status   string   `json:"status"`
}

func NewCompany(
	name string,
	document string,
	address *Address,
	status string,
) *Company {
	id := utils.NewUUID()
	return &Company{
		ID:       id,
		Name:     name,
		Document: document,
		Address:  address,
		Status:   status,
	}
}
