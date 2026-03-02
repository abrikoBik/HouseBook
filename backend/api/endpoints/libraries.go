package endpoints

import (
	"housebook/api/handlers"
	"housebook/api/middleware"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeLibraries(router chi.Router) {
	router.Route(consts.LibrariesEndpoint, func(r chi.Router) {
		r.With(middleware.JSONResponse).Get("/", handlers.ListLibraries)
		r.With(middleware.JSONResponse).Post("/", handlers.CreateLibrary)
		r.Route("/{library-uuid}", func(r chi.Router) {
			r.With(middleware.JSONResponse).Get("/", handlers.GetLibrary)
			r.Delete("/", handlers.DeleteLibrary)
			r.Route("/books", func(r chi.Router) {
				r.With(middleware.JSONResponse).Get("/", handlers.ListLibraryBooks)
				r.Patch("/", handlers.AddBookToLibrary)
				r.Delete("/{book-uuid}", handlers.DeleteBookFromLibrary)
			})
		})
	})
}
