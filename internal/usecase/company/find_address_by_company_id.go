package usecase

import (
	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
)

type FindAddressByCompanyIdUseCase struct {
	CompanyRepository entity.CompanyRepository
}

func NewFindAddressByCompanyIdUseCase(addressRepository entity.CompanyRepository) *FindAddressByCompanyIdUseCase {
	return &FindAddressByCompanyIdUseCase{
		CompanyRepository: addressRepository,
	}
}

func (uc FindAddressByCompanyIdUseCase) Execute(companyId string) (*dto.AddressResponseDto, error) {
	address, err := uc.CompanyRepository.FindAddressByCompanyId(companyId)
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
