package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type DeleteLevelUseCase struct {
	Repository entity.LevelRepository
}

func NewDeleteLevelUseCase(levelRepository entity.LevelRepository) *DeleteLevelUseCase {
	return &DeleteLevelUseCase{
		Repository: levelRepository,
	}
}

func (uc DeleteLevelUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
