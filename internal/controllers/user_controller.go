package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/dedicio/sisgares-accounts-service/internal/dto"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	usecase "github.com/dedicio/sisgares-accounts-service/internal/usecase/user"
	httpResponsePkg "github.com/dedicio/sisgares-accounts-service/pkg/response"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

type UserController struct {
	Repository entity.UserRepository
}

func NewUserController(userRepository entity.UserRepository) *UserController {
	return &UserController{
		Repository: userRepository,
	}
}

func (uc *UserController) FindAll(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	users, err := usecase.NewListUsersUseCase(uc.Repository).Execute(companyID)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewUsersResponse(users))
}

func (uc *UserController) FindById(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	user, err := usecase.NewFindUserByIdUseCase(uc.Repository).Execute(userId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Usuário"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewUserResponse(user))
}

func (uc *UserController) Create(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	payload := json.NewDecoder(r.Body)
	user := dto.UserDto{}
	err := payload.Decode(&user)
	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	user.CompanyId = companyID
	userSaved, err := usecase.NewCreateUserUseCase(uc.Repository).Execute(user)

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

func (uc *UserController) Delete(w http.ResponseWriter, r *http.Request) {
	userId := chi.URLParam(r, "id")
	err := usecase.NewDeleteUserUseCase(uc.Repository).Execute(userId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}

func (uc *UserController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	user := dto.UserDto{}
	err := payload.Decode(&user)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateUserUseCase(uc.Repository).Execute(user)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.UserResponseDto{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		Phone:   user.Phone,
		LevelId: user.LevelId,
	}

	render.Render(w, r, httpResponsePkg.NewUserResponse(output))
}
