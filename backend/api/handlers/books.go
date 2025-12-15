package handlers

import (
	"database/sql"
	"errors"
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
	book.UUID = *id
	if err := book.Select(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			render.Render(w, r, ErrorNotFound(err))
			return
		}
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
	var (
		book = &models.Book{}
		data = &models.BookRequest{}
	)
	id, err := util.PathUUID(consts.UUIDPathValue, r)
	if err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	book.UUID = *id
	if err := book.Select(); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	if err := book.Update(*data); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			render.Render(w, r, ErrorNotFound(err))
			return
		}
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.Render(w, r, bookResponse(book)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	var book models.Book
	id, err := util.PathUUID(consts.UUIDPathValue, r)
	if err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	book.UUID = *id
	if err := book.Delete(); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SearchBook(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
}
