package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type UpdateCompanyUseCase struct {
	Repository entity.CompanyRepository
}

func NewUpdateCompanyUseCase(companyRepository entity.CompanyRepository) *UpdateCompanyUseCase {
	return &UpdateCompanyUseCase{
		Repository: companyRepository,
	}
}

func (uc UpdateCompanyUseCase) Execute(input dto.CompanyDto) error {
	company, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	company.Name = input.Name
	company.Document = input.Document
	company.Status = input.Status
	company.Address.Street = input.Address.Street
	company.Address.Number = input.Address.Number
	company.Address.Complement = input.Address.Complement
	company.Address.Neighborhood = input.Address.Neighborhood
	company.Address.City = input.Address.City
	company.Address.State = input.Address.State
	company.Address.Country = input.Address.Country
	company.Address.ZipCode = input.Address.ZipCode

	err = uc.Repository.Update(company)
	if err != nil {
		return err
	}

	return nil
}
