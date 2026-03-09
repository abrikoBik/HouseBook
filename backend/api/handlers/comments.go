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

// CreateComment godoc
//
//	@summary		Создать отзыв
//	@description	Создаёт новый отзыв о книге.
//	@accept			json
//	@produce		json
//
//	@tags			Отзывы
//
//	@param			userid	body		int		true	"ID автора отзыва"
//	@param			bookid	body		int		true	"ID книги"
//	@param			body	body		string	true	"Текст отзыва"
//
//	@success		201		{object}	db.Comment
//	@failure		400		{object}	ErrorResponse
//	@failure		409		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/comments [post]
func CreateComment(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.CreateCommentParams
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
	comment, err := q.CreateComment(database.Context, data)
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
	Logger.Info().Interface("id", comment.ID).Msg("comment created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(comment); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// GetComment godoc
//
//	@summary		Получить отзыв
//	@description	Возвращает информацию об отзыве.
//	@accept			plain
//	@produce		json
//
//	@tags			Отзывы
//
//	@param			comment-id	path		pgtype.UUID	true	"UUID отзыва"
//
//	@success		200			{object}	db.Comment
//	@failure		400			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/comments/{comment-id} [get]
func GetComment(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.CommentIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.CommentIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	comment, err := q.GetComment(database.Context, uuid)
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
	Logger.Info().Interface("id", comment.ID).Msg("comment retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(comment); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// DeleteComment godoc
//
//	@summary		Удалить отзыв
//	@description	Удаляет отзыв из базы данных.
//	@accept			plain
//	@produce		json
//
//	@tags			Отзывы
//
//	@param			comment-id	path		pgtype.UUID	true	"UUID отзыва"
//
//	@success		200			{object}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/comments/{comment-id} [delete]
func DeleteComment(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.CommentIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.CommentIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteComment(database.Context, uuid); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", uuid).Msg("comment deleted")

	writer.WriteHeader(http.StatusOK)
}

// ListBookComments godoc
//
//	@summary		Список отзывов о книге
//	@description	Возвращает список всех отзывов о книге, хранящихся на сервере.
//	@accept			plain
//	@produce		json
//
//	@tags			Отзывы,Книги
//
//	@param			book-id	path		pgtype.UUID	true	"UUID книги"
//	@param			limit	query		int			false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int			false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.Comment
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/books/{book-id}/comments [get]
func ListBookComments(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListBookCommentsParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	if params.Bookid, err = util.PathUUID(consts.BookIDPathValue, request); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.BookIDPathValue))
		return
	}
	if params.Limit, params.Offset, err = util.GetSearchPagination(query); err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	comments, err := q.ListBookComments(database.Context, params)
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
	Logger.Info().Int("len(comments)", len(comments)).Msg("comments listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(comments); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// ListUserComments godoc
//
//	@summary		Список отзывов пользователя
//	@description	Возвращает список всех отзывов, написанных пользователем.
//	@accept			plain
//	@produce		json
//
//	@tags			Отзывы,Пользователи
//
//	@param			user-id	path		int	true	"ID пользователя"
//	@param			limit	query		int	false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int	false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.Comment
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/users/{user-id}/comments [get]
func ListUserComments(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListUserCommentsParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	if params.Userid, err = util.PathUUID(consts.UserIDPathValue, request); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.UserIDPathValue))
		return
	}
	if params.Limit, params.Offset, err = util.GetSearchPagination(query); err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	comments, err := q.ListUserComments(database.Context, params)
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
	Logger.Info().Int("len(comments)", len(comments)).Msg("comments listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(comments); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}
