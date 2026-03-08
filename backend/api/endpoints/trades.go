package endpoints

import (
	"housebook/api/handlers"
	"housebook/api/middleware"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeTrades(router chi.Router) {
	router.Route(consts.TradesEndpoint, func(r chi.Router) {
		r.With(middleware.JSONResponse).Get("/", handlers.ListTrades)
		r.With(middleware.JSONResponse).Post("/", handlers.CreateTrade)
		r.Route(consts.TradeIDPath, func(r chi.Router) {
			r.With(middleware.JSONResponse).Get("/", handlers.GetTrade)
			r.With(middleware.JSONResponse).Patch("/", handlers.ConfirmTrade)
			r.Delete("/", handlers.DeleteTrade)
			r.Route(consts.BooksEndpoint, func(r chi.Router) {
				r.With(middleware.JSONResponse).Get("/", handlers.ListTradeBooks)
				r.Patch("/", handlers.AddBookToTrade)
				r.Delete(consts.BookIDPath, handlers.DeleteTradeBook)
			})
		})
	})
}
