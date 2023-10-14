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

func NewLevelController(levelRepository entity.LevelRepository) *LevelController {
	return &LevelController{
		Repository: levelRepository,
	}
}

func (lc *LevelController) FindAll(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	levels, err := usecase.NewListLevelsUseCase(lc.Repository).Execute(companyID)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	render.Render(w, r, httpResponsePkg.NewLevelsResponse(levels))
}

func (lc *LevelController) FindById(w http.ResponseWriter, r *http.Request) {
	levelId := chi.URLParam(r, "id")
	level, err := usecase.NewFindLevelByIdUseCase(lc.Repository).Execute(levelId)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrNotFound(err, "Nível"))
		return
	}

	render.Render(w, r, httpResponsePkg.NewLevelResponse(level))
}

func (lc *LevelController) Create(w http.ResponseWriter, r *http.Request) {
	companyID := r.Header.Get("X-Company-ID")
	payload := json.NewDecoder(r.Body)
	level := dto.LevelDto{}
	err := payload.Decode(&level)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	level.CompanyId = companyID
	levelSaved, err := usecase.NewCreateLevelUseCase(lc.Repository).Execute(level)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.LevelResponseDto{
		ID:   levelSaved.ID,
		Name: levelSaved.Name,
	}

	render.Render(w, r, httpResponsePkg.NewLevelResponse(output))
}

func (lc *LevelController) Update(w http.ResponseWriter, r *http.Request) {
	payload := json.NewDecoder(r.Body)
	level := dto.LevelDto{}
	err := payload.Decode(&level)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInvalidRequest(err))
		return
	}

	err = usecase.NewUpdateLevelUseCase(lc.Repository).Execute(level)

	if err != nil {
		render.Render(w, r, httpResponsePkg.ErrInternalServerError(err))
		return
	}

	output := &dto.LevelResponseDto{
		ID:   level.ID,
		Name: level.Name,
	}

	render.Render(w, r, httpResponsePkg.NewLevelResponse(output))
}

func (lc *LevelController) Delete(w http.ResponseWriter, r *http.Request) {
	levelId := chi.URLParam(r, "id")
	err := usecase.NewDeleteLevelUseCase(lc.Repository).Execute(levelId)

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
