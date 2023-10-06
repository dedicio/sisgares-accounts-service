package response

import (
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
)

type LoginResponse struct {
	*dto.LoginResponseDto
}

func (pr *LoginResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewLoginResponse(login *dto.LoginResponseDto) *LoginResponse {
	return &LoginResponse{login}
}
