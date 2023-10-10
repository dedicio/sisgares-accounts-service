package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type ListLevelsUseCase struct {
	Repository entity.LevelRepository
}

func NewListLevelsUseCase(levelRepository entity.LevelRepository) *ListLevelsUseCase {
	return &ListLevelsUseCase{
		Repository: levelRepository,
	}
}

func (uc ListLevelsUseCase) Execute(companyID string) ([]*dto.LevelResponseDto, error) {
	levels, err := uc.Repository.FindAll(companyID)
	if err != nil {
		return nil, err
	}

	var output []*dto.LevelResponseDto
	for _, level := range levels {
		output = append(output, &dto.LevelResponseDto{
			ID:          level.ID,
			Name:        level.Name,
			Permissions: level.Permissions,
		})
	}

	return output, nil
}
