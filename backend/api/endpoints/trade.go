package endpoints

import (
	"housebook/api/handlers"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeTrades(router chi.Router) {
	router.Route(consts.TradesEndpoint, func(r chi.Router) {
		r.Get("/", handlers.ListTrades)
		r.Post("/", handlers.CreateTrade)
		r.Get("/{id}", handlers.GetTrade)
		r.Delete("/{id}", handlers.DeleteUser)
		r.Get("/search", handlers.SearchTrade)
	})
}
