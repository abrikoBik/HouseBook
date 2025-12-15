package handlers

import (
	"housebook/consts"
	"housebook/database/models"
	"housebook/util"
	"net/http"

	"github.com/go-chi/render"
)

func userResponse(user *models.User) *models.UserResponse {
	return &models.UserResponse{User: user}
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	var (
		user = &models.User{}
		data = &models.UserRequest{}
	)
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	if err := user.Insert(*data); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusCreated)
	if err := render.Render(w, r, userResponse(user)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	var (
		user    models.User
		id, err = util.PathID(consts.IDPathValue, r)
	)
	if err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	user.ID = id
	if err := user.Select(); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.Render(w, r, userResponse(&user)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func UpdateUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
