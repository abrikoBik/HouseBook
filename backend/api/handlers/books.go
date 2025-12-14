package handlers

import (
	"housebook/consts"
	"housebook/database/models"
	"housebook/util"
	"net/http"

	"github.com/go-chi/render"
)

func bookResponse(book *models.Book) *models.BookResponse {
	return &models.BookResponse{Book: book}
}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var (
		book = &models.Book{}
		data = &models.BookRequest{}
	)
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	if err := book.Insert(*data); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusCreated)
	if err := render.Render(w, r, bookResponse(book)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	id, err := util.PathUUID(consts.UUIDPathValue, r)
	if err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	if err := book.Select(*id); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.Render(w, r, bookResponse(&book)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
