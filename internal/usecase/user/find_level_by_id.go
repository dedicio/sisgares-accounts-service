package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type FindLevelByIdUseCase struct {
	Repository entity.LevelRepository
}

func NewFindLevelByIdUseCase(levelRepository entity.LevelRepository) *FindLevelByIdUseCase {
	return &FindLevelByIdUseCase{
		Repository: levelRepository,
	}
}

func (uc FindLevelByIdUseCase) Execute(id string) (*dto.LevelResponseDto, error) {
	level, err := uc.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	output := &dto.LevelResponseDto{
		ID:          level.ID,
		Name:        level.Name,
		Permissions: level.Permissions,
	}

	return output, nil
}
