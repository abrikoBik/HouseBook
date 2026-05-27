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
		r.Route(consts.LibraryIDPath, func(r chi.Router) {
			r.With(middleware.JSONResponse).Get("/", handlers.GetLibrary)
			r.Delete("/", handlers.DeleteLibrary)
			r.Route(consts.BooksEndpoint, func(r chi.Router) {
				r.With(middleware.JSONResponse).Get("/", handlers.ListLibraryBooks)
				r.Patch("/", handlers.AddBookToLibrary)
				r.Route(consts.BookIDPath, func(r chi.Router) {
					r.Patch("/", handlers.ChangeBookExchangeable)
					r.Delete("/", handlers.DeleteBookFromLibrary)
					r.Patch("/move", handlers.MoveBooksBetweenLibraries)
				})
			})
		})
	})
}
