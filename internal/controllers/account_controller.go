package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	companyUsecase "github.com/dedicio/sisgares-accounts-service/internal/usecase/company"
	usecase "github.com/dedicio/sisgares-accounts-service/internal/usecase/user"
	userUsecase "github.com/dedicio/sisgares-accounts-service/internal/usecase/user"
	httpResponsePkg "github.com/dedicio/sisgares-accounts-service/pkg/response"
	"github.com/go-chi/render"
)

type AccountController struct {
	CompanyRepository entity.CompanyRepository
	LevelRepository   entity.LevelRepository
	UserRepository    entity.UserRepository
}

func NewAccountController(
	companyRepository entity.CompanyRepository,
	levelRepository entity.LevelRepository,
	userRepository entity.UserRepository,
) *AccountController {
	return &AccountController{
		CompanyRepository: companyRepository,
		LevelRepository:   levelRepository,
		UserRepository:    userRepository,
	}
}

func (ac *AccountController) CreateAccount(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	account := dto.AccountInputDto{}
	err := payload.Decode(&account)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	company := entity.NewCompany(
		account.Company,
		"",
		&entity.Address{},
		"active",
	)
	companyDto := dto.CompanyDto{
		ID:       company.ID,
		Name:     company.Name,
		Document: company.Document,
		Address: dto.CompanyAddressDto{
			Street:       company.Address.Street,
			Number:       company.Address.Number,
			Complement:   company.Address.Complement,
			Neighborhood: company.Address.Neighborhood,
			City:         company.Address.City,
			State:        company.Address.State,
			Country:      company.Address.Country,
			ZipCode:      company.Address.ZipCode,
		},
		Status: company.Status,
	}

	companySaved, err := companyUsecase.NewCreateCompanyUseCase(ac.CompanyRepository).Execute(companyDto)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	level := entity.NewLevel(
		"Administrador",
		companySaved.ID,
		[]string{},
	)
	levelDto := dto.LevelDto{
		ID:          level.ID,
		Name:        level.Name,
		Permissions: level.Permissions,
		CompanyId:   level.CompanyId,
	}

	levelSaved, err := userUsecase.NewCreateLevelUseCase(ac.LevelRepository).Execute(levelDto)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	user := entity.NewUser(
		account.Name,
		account.Email,
		"",
		account.Password,
		levelSaved.ID,
		companySaved.ID,
	)
	userDto := dto.UserDto{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		Phone:     user.Phone,
		Password:  user.Password,
		LevelId:   user.LevelId,
		CompanyId: user.CompanyId,
	}

	userSaved, err := userUsecase.NewCreateUserUseCase(ac.UserRepository).Execute(userDto)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.UserResponseDto{
		ID:      userSaved.ID,
		Name:    userSaved.Name,
		Email:   userSaved.Email,
		Phone:   userSaved.Phone,
		LevelId: userSaved.LevelId,
	}
	render.Render(w, r, httpResponsePkg.NewUserResponse(output))
}

func (ac *AccountController) Login(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	login := dto.LoginDto{}
	err := payload.Decode(&login)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	userLogged, err := usecase.NewLoginUseCase(ac.UserRepository).Execute(login)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewLoginResponse(userLogged))
}
