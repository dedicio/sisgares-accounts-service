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

type CompanyController struct {
	CompanyRepository entity.CompanyRepository
	AddressRepository entity.AddressRepository
}

func NewCompanyController(
	companyRepository entity.CompanyRepository,
	addressRepository entity.AddressRepository,
) *CompanyController {
	return &CompanyController{
		CompanyRepository: companyRepository,
		AddressRepository: addressRepository,
	}
}

func (cc *CompanyController) FindAll(w http.ResponseWriter, r *http.Request) {
	categories, err := usecase.NewListCompaniesUseCase(cc.CompanyRepository).Execute()

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewCompaniesResponse(categories))
}

func (cc *CompanyController) FindById(w http.ResponseWriter, r *http.Request) {
	companyId := chi.URLParam(r, "id")
	company, err := usecase.NewFindCompanyByIdUseCase(cc.CompanyRepository).Execute(companyId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Categoria"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewCompanyResponse(company))
}

func (cc *CompanyController) Create(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	company := dto.CompanyDto{}
	err := payload.Decode(&company)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	companySaved, err := usecase.NewCreateCompanyUseCase(cc.CompanyRepository).Execute(company)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	address := dto.AddressDto{
		Street:       company.Address.Street,
		Number:       company.Address.Number,
		Complement:   company.Address.Complement,
		Neighborhood: company.Address.Neighborhood,
		City:         company.Address.City,
		State:        company.Address.State,
		Country:      company.Address.Country,
		ZipCode:      company.Address.ZipCode,
		CompanyId:    companySaved.ID,
	}

	addressSaved, err := usecase.NewCreateAddressUseCase(cc.AddressRepository).Execute(address)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.CompanyDto{
		ID:       companySaved.ID,
		Name:     companySaved.Name,
		Document: companySaved.Document,
		Address: dto.CompanyAddressDto{
			Street:       addressSaved.Street,
			Number:       addressSaved.Number,
			Complement:   addressSaved.Complement,
			Neighborhood: addressSaved.Neighborhood,
			City:         addressSaved.City,
			State:        addressSaved.State,
			Country:      addressSaved.Country,
			ZipCode:      addressSaved.ZipCode,
		},
		Status: companySaved.Status,
	}
	render.Render(w, r, httpResponsePkg.NewCompanyResponse(output))
}

func (cc *CompanyController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	company := dto.CompanyDto{}
	err := payload.Decode(&company)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateCompanyUseCase(cc.CompanyRepository).Execute(company)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	address := dto.AddressDto{
		Street:       company.Address.Street,
		Number:       company.Address.Number,
		Complement:   company.Address.Complement,
		Neighborhood: company.Address.Neighborhood,
		City:         company.Address.City,
		State:        company.Address.State,
		Country:      company.Address.Country,
		ZipCode:      company.Address.ZipCode,
		CompanyId:    company.ID,
	}

	err = usecase.NewUpdateAddressUseCase(cc.AddressRepository).Execute(address)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.CompanyDto{
		ID:       company.ID,
		Name:     company.Name,
		Document: company.Document,
		Address: dto.CompanyAddressDto{
			Street:       address.Street,
			Number:       address.Number,
			Complement:   address.Complement,
			Neighborhood: address.Neighborhood,
			City:         address.City,
			State:        address.State,
			Country:      address.Country,
			ZipCode:      address.ZipCode,
		},
		Status: company.Status,
	}
	render.Render(w, r, httpResponsePkg.NewCompanyResponse(output))
}

func (cc *CompanyController) Delete(w http.ResponseWriter, r *http.Request) {
	companyId := chi.URLParam(r, "id")
	err := usecase.NewDeleteCompanyUseCase(cc.CompanyRepository).Execute(companyId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	err = usecase.NewDeleteAddressUseCase(cc.AddressRepository).Execute(companyId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}
