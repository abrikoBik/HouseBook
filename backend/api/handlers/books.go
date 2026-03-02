package handlers

import (
	"encoding/json"
	"errors"
	"housebook/consts"
	"housebook/database"
	"housebook/database/db"
	"housebook/util"
	"net/http"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/unrolled/render"
)

// CreateBook godoc
//
// @summary Создать книгу
// @description Создаёт новую книгу.
// @accept json
// @produce json
//
// @param author body string true "Полное имя автора"
// @param cover_type body string false "Тип обложки (твёрдая, мягкая)"
// @param description body string false "Описание книги"
// @param genre body string true "Жанр книги"
// @param isbn body string true "ISBN книги"
// @param language body string true "Язык книги"
// @param page_count body int true "Количество страниц в книге"
// @param publisher body string true "Издательство книги"
// @param title body string true "Название книги"
// @param year body int true "Год выпуска книги"
//
// @failure 401 {object} ErrorResponse
// @failure 409 {object} ErrorResponse
// @failure 500 {object} ErrorResponse
// @router /api/books [post]
func CreateBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.CreateBookParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	book, err := q.CreateBook(database.Context, data)
	if err != nil {
		var (
			pgErr   *pgconn.PgError
			toWrite ErrorResponse
		)
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			toWrite = ErrorDuplicate(err, pgErr)
		} else {
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
	uuid, err := util.PathUUID(consts.UUIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.UUIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	book, err := q.GetBook(database.Context, uuid)
	if err != nil {
		var toWrite ErrorResponse
		if errors.Is(err, pgx.ErrNoRows) {
			toWrite = ErrorNotFound(err)
		} else {
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
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	book, err := q.UpdateBook(database.Context, data)
	if err != nil {
		var (
			pgErr   *pgconn.PgError
			toWrite ErrorResponse
		)
		errors.As(err, &pgErr)
		if pgErr.Code == "23505" {
			toWrite = ErrorDuplicate(err, pgErr)
		} else {
			toWrite = ErrorGeneric(err)
		}
		RenderError(r, writer, toWrite)
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
		RenderError(r, writer, ErrorBadID(err, consts.UUIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteBook(database.Context, uuid); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
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
	params, err = util.GetSearchParams(query)
	if err != nil {
		RenderError(r, writer, ErrorBadSearchParams(err, "book"))
		return
	}
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.SearchBooks(database.Context, params)
	if err != nil {
		var toWrite ErrorResponse
		if errors.Is(err, pgx.ErrNoRows) {
			toWrite = ErrorNotFound(err)
		} else {
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
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.ListBooks(database.Context, params)
	if err != nil {
		var toWrite ErrorResponse
		if errors.Is(err, pgx.ErrNoRows) {
			toWrite = ErrorNotFound(err)
		} else {
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
