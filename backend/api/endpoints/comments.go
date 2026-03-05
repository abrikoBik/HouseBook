package endpoints

import (
	"housebook/api/handlers"
	"housebook/api/middleware"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeComments(router chi.Router) {
	router.Route(consts.CommentsEndpoint, func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.CreateComment)
		r.Route("/{uuid}", func(r chi.Router) {
			r.With(middleware.JSONResponse).Get("/", handlers.GetComment)
			r.Delete("/", handlers.DeleteComment)
		})
	})
}
