package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type DeleteAddressByCompanyIdUseCase struct {
	Repository entity.AddressRepository
}

func NewDeleteAddressByCompanyIdUseCase(addressRepository entity.AddressRepository) *DeleteAddressByCompanyIdUseCase {
	return &DeleteAddressByCompanyIdUseCase{
		Repository: addressRepository,
	}
}

func (uc DeleteAddressByCompanyIdUseCase) Execute(companyId string) error {
	err := uc.Repository.DeleteByCompanyId(companyId)
	if err != nil {
		return err
	}

	return nil
}
