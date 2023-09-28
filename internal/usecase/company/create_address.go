package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type CreateAddressUseCase struct {
	Repository entity.AddressRepository
}

func NewCreateAddressUseCase(addressRepository entity.AddressRepository) *CreateAddressUseCase {
	return &CreateAddressUseCase{
		Repository: addressRepository,
	}
}

func (uc CreateAddressUseCase) Execute(input dto.AddressDto) (*dto.AddressDto, error) {
	address := entity.NewAddress(
		input.Street,
		input.Number,
		input.Complement,
		input.Neighborhood,
		input.City,
		input.State,
		input.Country,
		input.ZipCode,
		input.CompanyId,
	)

	err := uc.Repository.Create(address)
	if err != nil {
		return nil, err
	}

	output := &dto.AddressDto{
		ID:           address.ID,
		Street:       address.Street,
		Number:       address.Number,
		Complement:   address.Complement,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		ZipCode:      address.ZipCode,
		CompanyId:    address.CompanyId,
	}

	return output, nil
}
