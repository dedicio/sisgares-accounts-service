package routes

import (
	"github.com/dedicio/sisgares-accounts-service/internal/controllers"
	"github.com/dedicio/sisgares-accounts-service/internal/entity"
	"github.com/go-chi/chi/v5"
)

type AccountsRoutes struct {
	Controller controllers.AccountController
}

func NewAccountsRoutes(
	companyRepository entity.CompanyRepository,
	levelRepository entity.LevelRepository,
	userRepository entity.UserRepository,
) *AccountsRoutes {
	return &AccountsRoutes{
		Controller: *controllers.NewAccountController(
			companyRepository,
			levelRepository,
			userRepository,
		),
	}
}

func (ar AccountsRoutes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Route("/", func(router chi.Router) {
		router.Post("/register", ar.Controller.CreateAccount)
		router.Post("/login", ar.Controller.Login)
	})

	return router
}
