package routes

import (
	"github.com/dedicio/sisgares-accounts-service/internal/controllers"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/go-chi/chi/v5"
)

type LevelRoutes struct {
	Controller controllers.LevelController
}

func NewLevelRoutes(repository entity.LevelRepository) *LevelRoutes {
	return &LevelRoutes{
		Controller: *controllers.NewLevelController(repository),
	}
}

func (lr LevelRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/", lr.Controller.FindAll)
		router.Post("/", lr.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", lr.Controller.FindById)
			router.Delete("/", lr.Controller.Delete)
			router.Put("/", lr.Controller.Update)
			router.Get("/users", lr.Controller.FindUsersByLevelId)
		})
	})

	return router
}
