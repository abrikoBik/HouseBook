package endpoints

import (
	"housebook/api/handlers"
	"housebook/api/middleware"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeTrades(router chi.Router) {
	router.Route(consts.TradesEndpoint, func(r chi.Router) {
		// TODO: nest
		r.With(middleware.JSONResponse).Get("/", handlers.ListTrades)
		r.With(middleware.JSONResponse).Post("/", handlers.CreateTrade)
		r.With(middleware.JSONResponse).Get("/{uuid}", handlers.GetTrade)
		r.With(middleware.JSONResponse).Patch("/{uuid}", handlers.ConfirmTrade)
		r.Delete("/{uuid}", handlers.DeleteTrade)

		r.With(middleware.JSONResponse).Get("/{uuid}/books", handlers.ListTradeBooks)
		r.Patch("/{uuid}/books", handlers.AddBookToTrade)
		r.Delete("/{trade-uuid}/books/{book-uuid}", handlers.DeleteTradeBook)
	})
}
