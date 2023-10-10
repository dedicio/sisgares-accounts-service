package response

import (
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
)

type UserResponse struct {
	*dto.UserResponseDto
}

func (pr *UserResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUserResponse(product *dto.UserResponseDto) *UserResponse {
	return &UserResponse{product}
}

type UsersResponse struct {
	Users []*dto.UserResponseDto `json:"items"`
}

func (pr *UsersResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewUsersResponse(products []*dto.UserResponseDto) *UsersResponse {
	return &UsersResponse{products}
}
