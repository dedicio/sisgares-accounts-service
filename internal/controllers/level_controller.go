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

type LevelController struct {
	Repository entity.LevelRepository
}

func NewLevelController(positionRepository entity.LevelRepository) *LevelController {
	return &LevelController{
		Repository: positionRepository,
	}
}

func (lc *LevelController) FindAll(w http.ResponseWriter, r *http.Request) {
	positions, err := usecase.NewListLevelsUseCase(lc.Repository).Execute()

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewLevelsResponse(positions))
}

func (lc *LevelController) FindById(w http.ResponseWriter, r *http.Request) {
	positionId := chi.URLParam(r, "id")
	position, err := usecase.NewFindLevelByIdUseCase(lc.Repository).Execute(positionId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Nível"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewLevelResponse(position))
}

func (lc *LevelController) Create(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	position := dto.LevelDto{}
	err := payload.Decode(&position)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	positionSaved, err := usecase.NewCreateLevelUseCase(lc.Repository).Execute(position)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.LevelResponseDto{
		ID:          positionSaved.ID,
		Name:        positionSaved.Name,
		Permissions: positionSaved.Permissions,
	}

	render.Render(w, r, httpResponsePkg.NewLevelResponse(output))
}

func (lc *LevelController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	position := dto.LevelDto{}
	err := payload.Decode(&position)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateLevelUseCase(lc.Repository).Execute(position)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.LevelResponseDto{
		ID:          position.ID,
		Name:        position.Name,
		Permissions: position.Permissions,
	}

	render.Render(w, r, httpResponsePkg.NewLevelResponse(output))
}

func (lc *LevelController) Delete(w http.ResponseWriter, r *http.Request) {
	positionId := chi.URLParam(r, "id")
	err := usecase.NewDeleteLevelUseCase(lc.Repository).Execute(positionId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, nil)
}

func (lc *LevelController) FindUsersByLevelId(w http.ResponseWriter, r *http.Request) {
	levelId := chi.URLParam(r, "id")
	users, err := usecase.NewListUsersByLevelUseCase(lc.Repository).Execute(levelId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewUsersResponse(users))
}
