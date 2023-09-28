package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type CreateCompanyUseCase struct {
	Repository entity.CompanyRepository
}

func NewCreateCompanyUseCase(companyRepository entity.CompanyRepository) *CreateCompanyUseCase {
	return &CreateCompanyUseCase{
		Repository: companyRepository,
	}
}

func (uc CreateCompanyUseCase) Execute(input dto.CompanyDto) (*dto.CompanyDto, error) {
	company := entity.NewCompany(
		input.Name,
		input.Document,
		&entity.Address{},
		input.Status,
	)

	err := uc.Repository.Create(company)
	if err != nil {
		return nil, err
	}

	output := &dto.CompanyDto{
		ID:       company.ID,
		Name:     company.Name,
		Document: company.Document,
		Status:   company.Status,
	}

	return output, nil
}
