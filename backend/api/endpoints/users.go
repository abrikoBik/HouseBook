package endpoints

import (
	"housebook/api/handlers"

	"github.com/go-chi/chi/v5"
)

const (
	UsersPattern = "/users"
)

func routeUsers(router chi.Router) {
	router.Route(UsersPattern, func(r chi.Router) {
		r.Post("/", handlers.CreateUser)
		r.Get("/{id}", handlers.GetUser)
		r.Patch("/{id}", handlers.UpdateUser)
		r.Delete("/{id}", handlers.DeleteUser)
	})
}
