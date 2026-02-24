package endpoints

import (
	"housebook/api/handlers"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeBooks(router chi.Router) {
	router.Route(consts.BooksEndpoint, func(r chi.Router) {
		r.Get("/", handlers.ListBooks)
		r.Post("/", handlers.CreateBook)
		r.Get("/{uuid}", handlers.GetBook)
		r.Patch("/{uuid}", handlers.UpdateBook)
		r.Delete("/{uuid}", handlers.DeleteBook)
		r.Get("/search", handlers.SearchBook)
	})
}
