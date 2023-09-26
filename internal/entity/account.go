package entity

import (
	"github.com/dedicio/sisgares-accounts-service/pkg/utils"
)

type AccountRepository interface {
	FindById(id string) (*Account, error)
	FindAll() ([]*Account, error)
	Create(account *Account) error
	Update(account *Account) error
	Delete(id string) error
}

type CategoryRepository interface {
	FindById(id string) (*Category, error)
	FindAll() ([]*Category, error)
	Create(category *Category) error
	Update(category *Category) error
	Delete(id string) error
	FindAccountsByCategoryId(categoryId string) ([]*Account, error)
}

type Category struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyId string `json:"company_id"`
}

func NewCategory(name string, companyId string) *Category {
	id := utils.NewUUID()
	return &Category{
		ID:        id,
		Name:      name,
		CompanyId: companyId,
	}
}

type Account struct {
	ID          string  `json:"id"`
	Name        string  `json:"name"`
	Description string  `json:"description"`
	Price       float64 `json:"price"`
	Image       string  `json:"image"`
	CategoryId  string  `json:"category_id"`
	CompanyId   string  `json:"company_id"`
}

func NewAccount(
	name string,
	description string,
	price float64,
	image string,
	categoryId string,
	companyId string,
) *Account {
	id := utils.NewUUID()
	return &Account{
		ID:          id,
		Name:        name,
		Description: description,
		Price:       price,
		Image:       image,
		CategoryId:  categoryId,
		CompanyId:   companyId,
	}
}
