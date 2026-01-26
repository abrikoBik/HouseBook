package endpoints

import (
	"housebook/api/handlers"

	"github.com/go-chi/chi/v5"
)

const (
	BooksPattern = "/books"
)

func routeBooks(router chi.Router) {
	router.Route(BooksPattern, func(r chi.Router) {
		r.Get("/", handlers.ListBooks)
		r.Post("/", handlers.CreateBook)
		r.Get("/{uuid}", handlers.GetBook)
		r.Patch("/{uuid}", handlers.UpdateBook)
		r.Delete("/{uuid}", handlers.DeleteBook)
		r.Get("/search", handlers.SearchBook)
	})
}
