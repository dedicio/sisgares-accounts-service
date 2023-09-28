package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type ListUsersUseCase struct {
	UserRepository entity.UserRepository
}

func NewListUsersUseCase(userRepository entity.UserRepository) *ListUsersUseCase {
	return &ListUsersUseCase{
		UserRepository: userRepository,
	}
}

func (uc ListUsersUseCase) Execute() ([]*dto.UserResponseDto, error) {
	users, err := uc.UserRepository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*dto.UserResponseDto
	for _, user := range users {
		output = append(output, &dto.UserResponseDto{
			ID:      user.ID,
			Name:    user.Name,
			Email:   user.Email,
			Phone:   user.Phone,
			LevelId: user.LevelId,
		})
	}

	return output, nil
}
