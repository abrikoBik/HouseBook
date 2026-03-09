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

// CreateLibrary godoc
//
//	@summary		Создать библиотеку
//	@description	Создаёт новую пользовательскую библиотеку.
//	@accept			json
//	@produce		json
//
//	@tags			Библиотеки
//
//	@param			userid	body		int		true	"ID владельца библиотеки"
//	@param			name	body		string	true	"Название библиотеки"	Example(Избранное)
//
//	@success		201		{object}	db.Library
//	@failure		400		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/libraries [post]
func CreateLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.CreateLibraryParams
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
	library, err := q.CreateLibrary(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", library.ID).Msg("library created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(library); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// GetLibrary godoc
//
//	@summary		Получить библиотеку
//	@description	Возвращает информацию о библиотеке.
//	@accept			plain
//	@produce		json
//
//	@tags			Библиотеки
//
//	@param			library-id	path		pgtype.UUID	true	"UUID библиотеки"
//
//	@success		200			{object}	db.Library
//	@failure		400			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/libraries/{library-id} [get]
func GetLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	libraryUUID, err := util.PathUUID(consts.LibraryIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.LibraryIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	library, err := q.GetLibrary(database.Context, libraryUUID)
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
	Logger.Info().Interface("id", library.ID).Msg("library retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(library); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// DeleteLibrary godoc
//
//	@summary		Удалить библиотеку
//	@description	Удаляет библиотеку из базы данных.
//	@accept			plain
//	@produce		json
//
//	@tags			Библиотеки
//
//	@param			library-id	path		pgtype.UUID	true	"UUID библиотеки"
//
//	@success		200			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/libraries/{library-id} [delete]
func DeleteLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	libraryUUID, err := util.PathUUID(consts.LibraryIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.LibraryIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteLibrary(database.Context, libraryUUID); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", libraryUUID).Msg("library deleted")

	writer.WriteHeader(http.StatusOK)
}

// ListLibraries godoc
//
//	@summary		Список библиотек
//	@description	Возвращает список всех библиотек, хранящихся на сервере.
//	@accept			plain
//	@produce		json
//
//	@tags			Библиотеки
//
//	@param			limit	query		int	false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int	false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.Library
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/libraries [get]
func ListLibraries(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListLibrariesParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	if params.Limit, params.Offset, err = util.GetSearchPagination(query); err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	libraries, err := q.ListLibraries(database.Context, params)
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
	Logger.Info().Int("len(libraries)", len(libraries)).Msg("libraries listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(libraries); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// ListUserLibraries godoc
//
//	@summary		Список библиотек пользователя
//	@description	Возвращает список всех библиотек, принадлежащих пользователю.
//	@accept			plain
//	@produce		json
//
//	@tags			Библиотеки,Пользователи
//
//	@param			user-id	path		int	true	"ID пользователя"
//	@param			limit	query		int	false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int	false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.Library
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/users/{user-id}/libraries [get]
func ListUserLibraries(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListUserLibrariesParams
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
	libraries, err := q.ListUserLibraries(database.Context, params)
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
	Logger.Info().
		Str("userID", params.Userid.String()).
		Int("len(libraries)", len(libraries)).
		Msg("libraries listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(libraries); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// AddBookToLibrary godoc
//
//	@summary		Добавить книгу
//	@description	Добавляет книгу в существующую библиотеку.
//	@accept			json
//	@produce		json
//
//	@tags			Библиотеки,Книги
//
//	@param			library-id	path		string		true	"UUID библиотеки"
//	@param			bookid		body		pgtype.UUID	true	"ID книги"
//	@param			quantity	body		int			false	"Количество экземпляров книги"
//
//	@success		201			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		409			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/libraries/{library-id}/books [patch]
func AddBookToLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.AddBookToLibraryParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}
	if data.Libraryid, err = util.PathUUID(consts.LibraryIDPathValue, request); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.LibraryIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err = q.AddBookToLibrary(database.Context, data); err != nil {
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
	Logger.Info().
		Interface("libraryID", data.Libraryid).
		Interface("bookID", data.Bookid).
		Msg("book added to library")

	writer.WriteHeader(http.StatusCreated)
}

// ListLibraryBooks godoc
//
//	@summary		Список книг в библиотеке
//	@description	Возвращает список всех книг, имеющихся в библиотеке.
//	@accept			plain
//	@produce		json
//
//	@tags			Библиотеки,Книги
//
//	@param			library-id	path		pgtype.UUID	true	"UUID библиотеки"
//	@param			limit		query		int			false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset		query		int			false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200			{array}		db.LibraryBook
//	@failure		400			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/libraries/{library-id}/books [get]
func ListLibraryBooks(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListLibraryBooksParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	if params.ID, err = util.PathUUID(consts.LibraryIDPathValue, request); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.LibraryIDPathValue))
		return
	}
	if params.Limit, params.Offset, err = util.GetSearchPagination(query); err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.ListLibraryBooks(database.Context, params)
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

// DeleteBookFromLibrary godoc
//
//	@summary		Удалить книгу
//	@description	Удаляет книгу из существующей библиотеки.
//	@accept			json
//	@produce		json
//
//	@tags			Библиотеки,Книги
//
//	@param			library-id	path		pgtype.UUID	true	"UUID библиотеки"
//	@param			book-id		path		pgtype.UUID	true	"UUID книги"
//
//	@success		201			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/libraries/{library-id}/books/{book-id} [delete]
func DeleteBookFromLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.DeleteBookFromLibraryParams
		r      = render.New()
		q      = db.New(database.Pool)
	)

	if params.Libraryid, err = util.PathUUID(consts.LibraryIDPath, request); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.LibraryIDPath))
		return
	}
	if params.Bookid, err = util.PathUUID(consts.BookIDPathValue, request); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.BookIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteBookFromLibrary(database.Context, params); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().
		Interface("libraryID", params.Libraryid).
		Interface("bookID", params.Bookid).
		Msg("book deleted from library")

	writer.WriteHeader(http.StatusOK)
}
