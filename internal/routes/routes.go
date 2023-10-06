package routes

import (
	"database/sql"

	"github.com/dedicio/sisgares-accounts-service/internal/infra/repository"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type Routes struct {
	DB *sql.DB
}

func NewRoutes(db *sql.DB) *Routes {
	return &Routes{
		DB: db,
	}
}

func (routes Routes) Routes() chi.Router {
	router := chi.NewRouter()

	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.URLFormat)
	router.Use(render.SetContentType(render.ContentTypeJSON))

	userRepository := repository.NewUserRepositoryPostgres(routes.DB)
	companyRepository := repository.NewCompanyRepositoryPostgres(routes.DB)
	levelRepository := repository.NewLevelRepositoryPostgres(routes.DB)
	addressRepository := repository.NewAddressRepositoryPostgres(routes.DB)

	router.Route("/v1", func(router chi.Router) {
		router.Mount("/users", NewUserRoutes(userRepository).Routes())
		router.Mount("/levels", NewLevelRoutes(levelRepository).Routes())
		router.Mount("/addresses", NewAddressRoutes(addressRepository).Routes())
		router.Mount("/companies", NewCompanyRoutes(
			companyRepository,
			addressRepository,
		).Routes())
	})

	// novas rotas abertas

	return router
}
