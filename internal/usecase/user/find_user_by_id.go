package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type FindUserByIdUseCase struct {
	UserRepository entity.UserRepository
}

func NewFindUserByIdUseCase(userRepository entity.UserRepository) *FindUserByIdUseCase {
	return &FindUserByIdUseCase{
		UserRepository: userRepository,
	}
}

func (uc FindUserByIdUseCase) Execute(id string) (*dto.UserResponseDto, error) {
	user, err := uc.UserRepository.FindById(id)

	if err != nil {
		return nil, err
	}

	return &dto.UserResponseDto{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		LevelId: user.LevelId,
	}, nil
}
