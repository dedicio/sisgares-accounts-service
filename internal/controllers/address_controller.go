package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	usecase "github.com/dedicio/sisgares-accounts-service/internal/usecase/company"
	httpResponsePkg "github.com/dedicio/sisgares-accounts-service/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type AddressController struct {
	Repository entity.AddressRepository
}

func NewAddressController(addressRepository entity.AddressRepository) *AddressController {
	return &AddressController{
		Repository: addressRepository,
	}
}

func (ac *AddressController) Create(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	address := dto.AddressDto{}
	err := payload.Decode(&address)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	addressSaved, err := usecase.NewCreateAddressUseCase(ac.Repository).Execute(address)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.AddressResponseDto{
		ID:           addressSaved.ID,
		Street:       addressSaved.Street,
		Number:       addressSaved.Number,
		Complement:   addressSaved.Complement,
		Neighborhood: addressSaved.Neighborhood,
		City:         addressSaved.City,
		State:        addressSaved.State,
		Country:      addressSaved.Country,
		ZipCode:      addressSaved.ZipCode,
	}

	render.Render(w, r, httpResponsePkg.NewAddressResponse(output))
}

func (ac *AddressController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	address := dto.AddressDto{}
	err := payload.Decode(&address)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateAddressUseCase(ac.Repository).Execute(address)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
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

	render.Render(w, r, httpResponsePkg.NewAddressResponse(output))
}

func (ac *AddressController) Delete(w http.ResponseWriter, r *http.Request) {
	addressId := chi.URLParam(r, "id")
	err := usecase.NewDeleteAddressUseCase(ac.Repository).Execute(addressId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}

func (ac *AddressController) FindAddressByCompanyId(w http.ResponseWriter, r *http.Request) {
	companyId := chi.URLParam(r, "id")
	address, err := usecase.NewFindAddressByCompanyIdUseCase(ac.Repository).Execute(companyId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
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

	render.Render(w, r, httpResponsePkg.NewAddressResponse(output))
}
