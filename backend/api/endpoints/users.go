package endpoints

import (
	"housebook/api/handlers"
	"housebook/api/middleware"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeUsers(router chi.Router) {
	router.Route(consts.UsersEndpoint, func(r chi.Router) {
		r.With(middleware.JSONResponse).Post("/", handlers.CreateUser)
		r.Route(consts.UserIDPath, func(r chi.Router) {
			r.With(middleware.JSONResponse).Get("/", handlers.GetUser)
			r.With(middleware.JSONResponse).Patch("/", handlers.UpdateUser)
			r.With(middleware.JSONResponse).Put("/image", handlers.SetProfileImage)
			r.Delete("/", handlers.DeleteUser)

			r.With(middleware.JSONResponse).Get(consts.CommentsEndpoint, handlers.ListUserComments)
			r.With(middleware.JSONResponse).Get(consts.TradesEndpoint, handlers.ListUserTrades)
			r.With(middleware.JSONResponse).Get(consts.LibrariesEndpoint, handlers.ListUserLibraries)
			r.With(middleware.JSONResponse).Get(consts.SessionsEndpoint, handlers.ListUserSessions)
		})
	})
}
