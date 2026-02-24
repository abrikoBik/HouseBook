package endpoints

import (
	"housebook/api/handlers"
	"housebook/consts"

	"github.com/go-chi/chi/v5"
)

func routeUsers(router chi.Router) {
	router.Route(consts.UsersEndpoint, func(r chi.Router) {
		r.Post("/", handlers.CreateUser)
		r.Get("/{id}", handlers.GetUser)
		r.Patch("/{id}", handlers.UpdateUser)
		r.Delete("/{id}", handlers.DeleteUser)
	})
}
