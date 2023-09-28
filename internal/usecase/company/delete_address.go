package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type DeleteAddressUseCase struct {
	Repository entity.AddressRepository
}

func NewDeleteAddressUseCase(addressRepository entity.AddressRepository) *DeleteAddressUseCase {
	return &DeleteAddressUseCase{
		Repository: addressRepository,
	}
}

func (uc DeleteAddressUseCase) Execute(id string) error {
	err := uc.Repository.Delete(id)
	if err != nil {
		return err
	}

	return nil
}
