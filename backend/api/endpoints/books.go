package endpoints

import (
	"housebook/api/handlers"
	"housebook/api/middleware"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeBooks(router chi.Router) {
	router.Route(consts.BooksEndpoint, func(r chi.Router) {
		r.With(middleware.JSONResponse).Get("/", handlers.ListBooks)
		r.With(middleware.JSONResponse).Post("/", handlers.CreateBook)
		r.With(middleware.JSONResponse).Get("/search", handlers.SearchBook)
		r.Route("/{uuid}", func(r chi.Router) {
			r.With(middleware.JSONResponse).Get("/", handlers.GetBook)
			r.With(middleware.JSONResponse).Patch("/", handlers.UpdateBook)
			r.Delete("/", handlers.DeleteBook)
			r.With(middleware.JSONResponse).Get("/comments", handlers.ListBookComments)
		})
	})
}
