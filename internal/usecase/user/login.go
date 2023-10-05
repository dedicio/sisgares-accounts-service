package usecase

import (
	"fmt"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/dedicio/sisgares-accounts-service/pkg/utils"
)

type LoginUseCase struct {
	UserRepository entity.UserRepository
}

func NewLoginUseCase(userRepository entity.UserRepository) *LoginUseCase {
	return &LoginUseCase{
		UserRepository: userRepository,
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

	fmt.Println("User: ", user, "enviou para o client")

	return &dto.LoginResponseDto{
		ID:        user.ID,
		CompanyID: user.CompanyId,
		Hash:      "string-aleatoria",
	}, nil
}
