package usecase

import "github.com/dedicio/sisgares-accounts-service/internal/entity"

type DeleteUserUseCase struct {
	Repository entity.UserRepository
}

func NewDeleteUserUseCase(userRepository entity.UserRepository) *DeleteUserUseCase {
	return &DeleteUserUseCase{
		Repository: userRepository,
	}
}

func (uc *DeleteUserUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
