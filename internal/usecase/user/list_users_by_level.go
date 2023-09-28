package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type ListUsersByLevelUseCase struct {
	LevelRepository entity.LevelRepository
}

func NewListUsersByLevelUseCase(userRepository entity.LevelRepository) *ListUsersByLevelUseCase {
	return &ListUsersByLevelUseCase{
		LevelRepository: userRepository,
	}
}

func (uc ListUsersByLevelUseCase) Execute(levelId string) ([]*dto.UserResponseDto, error) {
	users, err := uc.LevelRepository.FindUsersByLevelId(levelId)
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
