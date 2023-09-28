package response

import (
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
)

type AddressResponse struct {
	*dto.AddressResponseDto
}

func (cr *AddressResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewAddressResponse(category *dto.AddressResponseDto) *AddressResponse {
	return &AddressResponse{category}
}

type AddressesResponse struct {
	Addresses []*dto.AddressResponseDto
}

func (cr *AddressesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewAddressesResponse(categories []*dto.AddressResponseDto) *AddressesResponse {
	return &AddressesResponse{categories}
}
