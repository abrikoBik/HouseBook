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
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}
	data.Userid = userID

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
//	@success		200	{object}	db.Library
//	@failure		400	{object}	ErrorResponse
//	@failure		404	{object}	ErrorResponse
//	@failure		500	{object}	ErrorResponse
//	@router			/api/libraries/{uuid} [get]
func GetLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.GetLibraryParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	// TODO: clean up, match the docs
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	data.Userid = userID
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	data.ID = libraryUUID

	if PingDB(r, writer) != nil {
		return
	}
	library, err := q.GetLibrary(database.Context, data)
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
//	@success		200	{string}	string	""
//	@failure		400	{object}	ErrorResponse
//	@failure		500	{object}	ErrorResponse
//	@router			/api/libraries/{uuid} [delete]
func DeleteLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.DeleteLibraryParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	data.Userid = userID
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	data.ID = libraryUUID

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteLibrary(database.Context, data); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", libraryUUID).Msg("library deleted")

	writer.WriteHeader(http.StatusOK)
}

// ListLibraries godoc
//
//	@summary		Список библиотек
//	@description	Возвращает список всех библиотек, принадлежащих пользователю.
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
//	@router			/api/users/{id}/libraries [get]
func ListLibraries(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListLibrariesParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	params.Userid = userID
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
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

// AddBookToLibrary godoc
//
//	@summary		Добавить книгу
//	@description	Добавляет книгу в существующую библиотеку.
//	@accept			json
//	@produce		json
//
//	@tags			Библиотеки
//
//	@param			bookid		body		pgtype.UUID	true	"ID книги"
//	@param			quantity	body		int			false	"Количество экземпляров книги"
//
//	@success		201			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		409			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/libraries/{uuid}/books [patch]
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
	_, err = util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	data.Libraryid = libraryUUID

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
//	@tags			Библиотеки
//
//	@param			limit	query		int	false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int	false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.LibraryBook
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/libraries/{uuid}/books [get]
func ListLibraryBooks(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListLibraryBooksParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	userID, err := util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	params.Userid = userID
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	params.ID = libraryUUID
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
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
//	@tags			Библиотеки
//
//	@success		201	{string}	string	""
//	@failure		400	{object}	ErrorResponse
//	@failure		500	{object}	ErrorResponse
//	@router			/api/libraries/{library-uuid}/books/{book-uuid} [delete]
func DeleteBookFromLibrary(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.DeleteBookFromLibraryParams
		r      = render.New()
		q      = db.New(database.Pool)
	)

	_, err = util.PathID(consts.IDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.IDPathValue))
		return
	}
	libraryUUID, err := util.PathUUID("library-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "library-uuid"))
		return
	}
	params.Libraryid = libraryUUID
	bookID, err := util.PathUUID("book-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "book-uuid"))
		return
	}
	params.Bookid = bookID

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteBookFromLibrary(database.Context, params); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().
		Interface("libraryID", libraryUUID).
		Interface("bookID", bookID).
		Msg("book deleted from library")

	writer.WriteHeader(http.StatusOK)
}
