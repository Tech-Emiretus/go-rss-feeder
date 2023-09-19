package router

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func GetRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Use(middleware.Logger)
	registerRoutes(router)

	return router
}
