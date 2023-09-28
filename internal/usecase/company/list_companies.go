package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type ListCompanysUseCase struct {
	Repository entity.CompanyRepository
}

func NewListCompaniesUseCase(companyRepository entity.CompanyRepository) *ListCompanysUseCase {
	return &ListCompanysUseCase{
		Repository: companyRepository,
	}
}

func (uc ListCompanysUseCase) Execute() ([]*dto.CompanyDto, error) {
	companies, err := uc.Repository.FindAll()
	if err != nil {
		return nil, err
	}

	var output []*dto.CompanyDto
	for _, company := range companies {
		output = append(output, &dto.CompanyDto{
			ID:       company.ID,
			Name:     company.Name,
			Document: company.Document,
			Status:   company.Status,
			Address: dto.CompanyAddressDto{
				Street:       company.Address.Street,
				Number:       company.Address.Number,
				Complement:   company.Address.Complement,
				Neighborhood: company.Address.Neighborhood,
				City:         company.Address.City,
				State:        company.Address.State,
				Country:      company.Address.Country,
				ZipCode:      company.Address.ZipCode,
			},
		})
	}

	return output, nil
}
