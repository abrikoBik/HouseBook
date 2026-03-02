package handlers

import (
	"database/sql"
	"encoding/json"
	"errors"
	"housebook/consts"
	"housebook/database"
	"housebook/database/db"
	"housebook/util"
	"net/http"

	"github.com/unrolled/render"
)

func CreateBook(writer http.ResponseWriter, request *http.Request) {
	// TODO: create a small package for all rendering-related stuff
	var (
		err  error
		data db.CreateBookParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	book, err := q.CreateBook(database.Context, data)
	if err != nil {
		var toWrite ErrorResponse
		// TODO: handle more errors
		switch err {
		default:
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Interface("id", book.ID).Msg("book created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func GetBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	// TODO: commonize
	uuid, err := util.PathUUID(consts.UUIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	book, err := q.GetBook(database.Context, uuid)
	if err != nil {
		var toWrite ErrorResponse
		switch err {
		case sql.ErrNoRows:
			toWrite = ErrorNotFound(err)
		default:
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Interface("id", book.ID).Msg("book retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.UpdateBookParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	data.ID, err = util.PathUUID(consts.UUIDPathValue, request)
	if err = errors.Join(err, json.NewDecoder(request.Body).Decode(&data)); err != nil {
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	book, err := q.UpdateBook(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", book.ID).Msg("book updated")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.UUIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteBook(database.Context, uuid); err != nil {
		var toWrite ErrorResponse
		switch err {
		default:
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Interface("id", uuid).Msg("book deleted")

	writer.WriteHeader(http.StatusOK)
}

func SearchBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.SearchBooksParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	// TODO: eliminate doubled code
	params, err = util.GetSearchParams(query)
	if err != nil {
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.SearchBooks(database.Context, params)
	if err != nil {
		var toWrite ErrorResponse
		switch err {
		case sql.ErrNoRows:
			// TODO: decide if 404 is applicable here
			toWrite = ErrorNotFound(err)
		default:
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Int("len(books)", len(books)).Msg("books searched")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(books); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

func ListBooks(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListBooksParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
		RenderError(r, writer, ErrorBadRequest(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.ListBooks(database.Context, params)
	if err != nil {
		var toWrite ErrorResponse
		switch err {
		case sql.ErrNoRows:
			// TODO: decide if 404 is applicable here
			toWrite = ErrorNotFound(err)
		default:
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
		return
	}
	Logger.Info().Int("len(books)", len(books)).Msg("books listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(books); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}
