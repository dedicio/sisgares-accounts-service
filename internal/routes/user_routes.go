package routes

import (
	"github.com/dedicio/sisgares-accounts-service/internal/controllers"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/go-chi/chi/v5"
)

type UserRoutes struct {
	Controller controllers.UserController
}

func NewUserRoutes(repository entity.UserRepository) *UserRoutes {
	return &UserRoutes{
		Controller: *controllers.NewUserController(repository),
	}
}

func (ur UserRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/", ur.Controller.FindAll)
		router.Post("/", ur.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", ur.Controller.FindById)
			router.Delete("/", ur.Controller.Delete)
			router.Put("/", ur.Controller.Update)
		})
	})

	return router
}
