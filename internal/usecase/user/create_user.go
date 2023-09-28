package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type CreateUserUseCase struct {
	Repository entity.UserRepository
}

func NewCreateUserUseCase(userRepository entity.UserRepository) *CreateUserUseCase {
	return &CreateUserUseCase{
		Repository: userRepository,
	}
}

func (uc *CreateUserUseCase) Execute(input dto.UserDto) (*dto.UserDto, error) {
	user := entity.NewUser(
		input.Name,
		input.Email,
		input.Phone,
		input.Password,
		input.LevelId,
		input.CompanyId,
	)

	err := uc.Repository.Create(user)
	if err != nil {
		return nil, err
	}

	output := &dto.UserDto{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		LevelId: user.LevelId,
	}

	return output, nil
}
