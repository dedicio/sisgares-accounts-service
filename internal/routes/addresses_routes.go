package routes

import (
	"github.com/dedicio/sisgares-accounts-service/internal/controllers"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/go-chi/chi/v5"
)

type AddressRoutes struct {
	Controller controllers.AddressController
}

func NewAddressRoutes(repository entity.AddressRepository) *AddressRoutes {
	return &AddressRoutes{
		Controller: *controllers.NewAddressController(repository),
	}
}

func (ar AddressRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Post("/", ar.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Delete("/", ar.Controller.Delete)
			router.Put("/", ar.Controller.Update)
		})
	})

	return router
}
