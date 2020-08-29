package api

import (
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func routes(services *Services) *chi.Mux {

	r := chi.NewMux()

	r.Use(middleware.Recoverer)

	r.Get("/heartbeat", services.healthCheck)

	return r
}
