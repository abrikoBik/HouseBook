// Package endpoints contains all HTTP endpoints and their routes as strings.
package endpoints

import (
	"housebook/api/handlers"
	_ "housebook/docs"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func Register(router chi.Router) {
	router.Mount("/api", router)
	router.Get("/health", handlers.Health)
	routeUsers(router)
	routeBooks(router)
	routeTrades(router)
	router.Route("/docs", func(r chi.Router) {
		r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "/app/swagger.json")
		})
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL("/api/docs/swagger.json"),
		))
	})
}
