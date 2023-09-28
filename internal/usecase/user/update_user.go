package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type UpdateUserUseCase struct {
	Repository entity.UserRepository
}

func NewUpdateUserUseCase(userRepository entity.UserRepository) *UpdateUserUseCase {
	return &UpdateUserUseCase{
		Repository: userRepository,
	}
}

func (uc *UpdateUserUseCase) Execute(input dto.UserDto) error {
	user, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Phone = input.Phone
	user.LevelId = input.LevelId
	user.CompanyId = input.CompanyId

	err = uc.Repository.Update(user)
	if err != nil {
		return err
	}

	return nil
}
