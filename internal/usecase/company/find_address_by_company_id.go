package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type FindAddressByCompanyIdUseCase struct {
	AddressRepository entity.AddressRepository
}

func NewFindAddressByCompanyIdUseCase(addressRepository entity.AddressRepository) *FindAddressByCompanyIdUseCase {
	return &FindAddressByCompanyIdUseCase{
		AddressRepository: addressRepository,
	}
}

func (uc FindAddressByCompanyIdUseCase) Execute(companyId string) (*dto.AddressResponseDto, error) {
	address, err := uc.AddressRepository.FindAddressByCompanyId(companyId)
	if err != nil {
		return nil, err
	}

	output := &dto.AddressResponseDto{
		ID:           address.ID,
		Street:       address.Street,
		Number:       address.Number,
		Complement:   address.Complement,
		Neighborhood: address.Neighborhood,
		City:         address.City,
		State:        address.State,
		Country:      address.Country,
		ZipCode:      address.ZipCode,
	}

	return output, nil
}
