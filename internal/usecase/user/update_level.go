package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type UpdateLevelUseCase struct {
	Repository entity.LevelRepository
}

func NewUpdateLevelUseCase(levelRepository entity.LevelRepository) *UpdateLevelUseCase {
	return &UpdateLevelUseCase{
		Repository: levelRepository,
	}
}

func (uc UpdateLevelUseCase) Execute(input dto.LevelDto) error {
	level, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	level.Name = input.Name
	level.CompanyId = input.CompanyId

	err = uc.Repository.Update(level)
	if err != nil {
		return err
	}

	return nil
}
