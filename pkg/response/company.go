package response

import (
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
)

type CompanyResponse struct {
	*dto.CompanyDto
}

func (cr *CompanyResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewCompanyResponse(category *dto.CompanyDto) *CompanyResponse {
	return &CompanyResponse{category}
}

type CompaniesResponse struct {
	Companies []*dto.CompanyDto `json:"items"`
}

func (cr *CompaniesResponse) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func NewCompaniesResponse(categories []*dto.CompanyDto) *CompaniesResponse {
	return &CompaniesResponse{categories}
}
