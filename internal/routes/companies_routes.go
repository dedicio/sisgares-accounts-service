package routes

import (
	"github.com/dedicio/sisgares-accounts-service/internal/controllers"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/go-chi/chi/v5"
)

type CompanyRoutes struct {
	Controller controllers.CompanyController
}

func NewCompanyRoutes(
	companyRepository entity.CompanyRepository,
	addressRepository entity.AddressRepository,
) *CompanyRoutes {
	return &CompanyRoutes{
		Controller: *controllers.NewCompanyController(
			companyRepository,
			addressRepository,
		),
	}
}

func (cr CompanyRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Get("/", cr.Controller.FindAll)
		router.Post("/", cr.Controller.Create)

		router.Route("/{id}", func(router chi.Router) {
			router.Get("/", cr.Controller.FindById)
			router.Delete("/", cr.Controller.Delete)
			router.Put("/", cr.Controller.Update)
			router.Get("/addresses", cr.Controller.FindAddressByCompanyId)
		})
	})

	return router
}
