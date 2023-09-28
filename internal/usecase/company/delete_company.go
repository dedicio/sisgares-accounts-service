package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type DeleteCompanyUseCase struct {
	Repository entity.CompanyRepository
}

func NewDeleteCompanyUseCase(companyRepository entity.CompanyRepository) *DeleteCompanyUseCase {
	return &DeleteCompanyUseCase{
		Repository: companyRepository,
	}
}

func (uc DeleteCompanyUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
