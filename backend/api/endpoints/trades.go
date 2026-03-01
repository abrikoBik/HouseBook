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
		r.Delete("/{uuid}", handlers.DeleteTrade)

		r.Get("/{uuid}", handlers.GetTrade)

		r.Get("/{uuid}/books", handlers.ListTradeBooks)
		r.Patch("/{uuid}/books", handlers.AddBookToTrade)
		r.Delete("/{trade-uuid}/books/{book-uuid}", handlers.DeleteTradeBook)
	})
}
