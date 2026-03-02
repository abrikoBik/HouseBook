// Package endpoints contains all HTTP endpoints and their routes as strings.
package endpoints

import (
	"housebook/api/handlers"

	"github.com/go-chi/chi/v5"
)

func Register(router chi.Router) {
	router.Get("/health", handlers.Health)
	routeUsers(router)
	routeBooks(router)
	routeTrades(router)
}
