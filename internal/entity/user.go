package entity

import (
	"github.com/dedicio/sisgares-accounts-service/pkg/utils"
)

type UserRepository interface {
	FindById(id string) (*User, error)
	FindAll() ([]*User, error)
	Create(user *User) error
	Update(user *User) error
	Delete(id string) error
}

type LevelRepository interface {
	FindById(id string) (*Level, error)
	FindAll() ([]*Level, error)
	Create(level *Level) error
	Update(level *Level) error
	Delete(id string) error
	FindUsersByLevelId(levelId string) ([]*User, error)
}

type Level struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	AccountId string `json:"account_id"`
}

func NewLevel(name string, companyId string) *Level {
	id := utils.NewUUID()
	return &Level{
		ID:        id,
		Name:      name,
		AccountId: companyId,
	}
}

type User struct {
	ID        string  `json:"id"`
	Name      string  `json:"name"`
	Email     string  `json:"email"`
	Phone     float64 `json:"phone"`
	Password  string  `json:"password"`
	LevelId   string  `json:"level_id"`
	AccountId string  `json:"account_id"`
}

func NewUser(
	name string,
	email string,
	phone float64,
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
		AccountId: companyId,
	}
}
