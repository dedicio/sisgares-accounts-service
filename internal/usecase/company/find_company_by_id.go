package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type FindCompanyByIdUseCase struct {
	Repository entity.CompanyRepository
}

func NewFindCompanyByIdUseCase(companyRepository entity.CompanyRepository) *FindCompanyByIdUseCase {
	return &FindCompanyByIdUseCase{
		Repository: companyRepository,
	}
}

func (uc FindCompanyByIdUseCase) Execute(id string) (*dto.CompanyDto, error) {
	company, err := uc.Repository.FindById(id)
	if err != nil {
		return nil, err
	}

	output := &dto.CompanyDto{
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
	}

	return output, nil
}
