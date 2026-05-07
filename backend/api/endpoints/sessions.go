package endpoints

import (
	"housebook/api/handlers"
	"housebook/api/middleware"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeSessions(router chi.Router) {
	router.Route(consts.SessionsEndpoint, func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.CreateSession)
	})
}
