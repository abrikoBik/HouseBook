// Package endpoints contains all HTTP endpoints and their routes as strings.
package endpoints

import (
	"housebook/api/handlers"

	"github.com/go-chi/chi/v5"
)

func Register(router chi.Router) {
	router.Mount("/api", router)
	router.Get("/health", handlers.Health)
	routeDocs(router)
	routeUsers(router)
	routeBooks(router)
	routeTrades(router)
	routeLibraries(router)
	routeComments(router)
	routeSessions(router)
}
