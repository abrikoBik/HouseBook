package endpoints

import (
	"housebook/api/handlers"

	"github.com/go-chi/chi/v5"
)

const (
	TradesPattern = "/trades"
)

func routeTrades(router chi.Router) {
	router.Route(TradesPattern, func(r chi.Router) {
		r.Get("/", handlers.ListTrades)
		r.Post("/", handlers.CreateTrade)
		r.Get("/{id}", handlers.GetTrade)
		r.Delete("/{id}", handlers.DeleteUser)
		r.Get("/search", handlers.SearchTrade)
	})
}
