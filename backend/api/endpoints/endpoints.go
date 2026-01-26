// Package endpoints contains all HTTP endpoints and their routes as strings.
package endpoints

import (
	"housebook/api/handlers"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func Register(router chi.Router) {
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/health", handlers.Health)
	routeUsers(router)
	routeBooks(router)
	routeTrades(router)
}
