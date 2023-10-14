package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/dedicio/sisgares-accounts-service/pkg/utils"
)

type LoginUseCase struct {
	UserRepository  entity.UserRepository
	LevelRepository entity.LevelRepository
}

func NewLoginUseCase(
	userRepository entity.UserRepository,
	levelRepository entity.LevelRepository,
) *LoginUseCase {
	return &LoginUseCase{
		UserRepository:  userRepository,
		LevelRepository: levelRepository,
	}
}

func (uc LoginUseCase) Execute(login dto.LoginDto) (*dto.LoginResponseDto, error) {
	user, err := uc.UserRepository.FindByEmail(login.Email)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, err
	}

	validatePassword := utils.CheckPasswordHash(login.Password, user.Password)
	if !validatePassword {
		return nil, err
	}

	level, err := uc.LevelRepository.FindById(user.LevelId)
	if err != nil {
		return nil, err
	}

	return &dto.LoginResponseDto{
		ID:        user.ID,
		Name:      user.Name,
		Level:     level.Name,
		CompanyID: user.CompanyId,
	}, nil
}
