package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type UpdateAddressUseCase struct {
	Repository entity.AddressRepository
}

func NewUpdateAddressUseCase(addressRepository entity.AddressRepository) *UpdateAddressUseCase {
	return &UpdateAddressUseCase{
		Repository: addressRepository,
	}
}

func (uc UpdateAddressUseCase) Execute(input dto.AddressDto) error {
	address, err := uc.Repository.FindById(input.ID)
	if err != nil {
		return err
	}

	address.Street = input.Street
	address.Number = input.Number
	address.Complement = input.Complement
	address.Neighborhood = input.Neighborhood
	address.City = input.City
	address.State = input.State
	address.Country = input.Country
	address.ZipCode = input.ZipCode
	address.CompanyId = input.CompanyId

	err = uc.Repository.Update(address)
	if err != nil {
		return err
	}

	return nil
}
