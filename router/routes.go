package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/techemeritus/rss-feeder/helpers"
)

func registerRoutes(router *chi.Mux) {
	router.Get("/", func (response http.ResponseWriter, request *http.Request) {
		httpResponse := helpers.HttpResponse{Writer: response}
		httpResponse.Success("Rss Feed is Running...", 200)
	})

	router.Get("/error", func (response http.ResponseWriter, request *http.Request) {
		httpResponse := helpers.HttpResponse{Writer: response}
		httpResponse.Error("Rss Feed is Unhealthy :sad", 400)
	})
}
