package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type CreateLevelUseCase struct {
	Repository entity.LevelRepository
}

func NewCreateLevelUseCase(levelRepository entity.LevelRepository) *CreateLevelUseCase {
	return &CreateLevelUseCase{
		Repository: levelRepository,
	}
}

func (uc CreateLevelUseCase) Execute(input dto.LevelDto) (*dto.LevelDto, error) {
	level := entity.NewLevel(
		input.Name,
		input.CompanyId,
	)

	err := uc.Repository.Create(level)
	if err != nil {
		return nil, err
	}

	output := &dto.LevelDto{
		ID:   level.ID,
		Name: level.Name,
	}

	return output, nil
}
