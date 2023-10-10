package entity

import (
	"github.com/dedicio/sisgares-accounts-service/pkg/utils"
)

type UserRepository interface {
	FindById(id string) (*User, error)
	FindAll(companyID string) ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
	FindByEmail(email string) (*User, error)
}

type LevelRepository interface {
	FindById(id string) (*Level, error)
	FindAll(companyID string) ([]*Level, error)
	Create(level *Level) error
	Update(level *Level) error
	Delete(id string) error
	FindUsersByLevelId(levelId string) ([]*User, error)
}

type Permission struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	CompanyID string `json:"company_id"`
}

func NewPermission(name string, companyId string) *Permission {
	id := utils.NewUUID()
	return &Permission{
		ID:        id,
		Name:      name,
		CompanyID: companyId,
	}
}

type Level struct {
	ID          string   `json:"id"`
	Name        string   `json:"name"`
	CompanyId   string   `json:"company_id"`
	Permissions []string `json:"permissions"`
}

func NewLevel(
	name string,
	companyId string,
	permissions []string,
) *Level {
	id := utils.NewUUID()
	return &Level{
		ID:          id,
		Name:        name,
		CompanyId:   companyId,
		Permissions: permissions,
	}
}

type User struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Password  string `json:"password"`
	LevelId   string `json:"level_id"`
	CompanyId string `json:"company_id"`
}

func NewUser(
	name string,
	email string,
	phone string,
	password string,
	levelId string,
	companyId string,
) *User {
	id := utils.NewUUID()
	return &User{
		ID:        id,
		Name:      name,
		Email:     email,
		Phone:     phone,
		Password:  password,
		LevelId:   levelId,
		CompanyId: companyId,
	}
}
