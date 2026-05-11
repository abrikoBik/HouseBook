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
//	@summary		Создать книгу
//	@description	Создаёт новую книгу.
//	@accept			json
//	@produce		json
//
//	@tags			Книги
//
//	@param			author		body		string	true	"Полное имя автора"				Example(Известный автор)
//	@param			cover_type	body		string	false	"Тип обложки (твёрдая, мягкая)"	Example(твёрдая)
//	@param			description	body		string	false	"Описание книги"				Example(Крайне интересная книга об очень интересных вещах)
//	@param			genre		body		string	true	"Жанр книги"					Example(Роман)
//	@param			isbn		body		string	true	"ISBN книги"					Example(978-3-16-148410-0)
//	@param			language	body		string	false	"Язык книги"					Example(Русский)
//	@param			page_count	body		int		false	"Количество страниц в книге"	Example(100)
//	@param			publisher	body		string	true	"Издательство книги"			Example(Хороший Издатель)
//	@param			title		body		string	true	"Название книги"				Example(Интересная книга)
//	@param			year		body		int		false	"Год выпуска книги"				Example(2026)
//
//	@success		201			{object}	db.Book
//	@failure		400			{object}	ErrorResponse
//	@failure		409			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/books [post]
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

// GetBook godoc
//
//	@summary		Получить книгу
//	@description	Возвращает информацию о книге.
//	@accept			plain
//	@produce		json
//
//	@tags			Книги
//
//	@param			book-id	path		pgtype.UUID	true	"UUID книги"
//
//	@success		200		{object}	db.Book
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/books/{book-id} [get]
func GetBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.BookIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.BookIDPathValue))
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

// UpdateBook godoc
//
//	@summary		Изменить книгу
//	@description	Изменяет информацию о существующей книге.
//	@accept			json
//	@produce		json
//
//	@tags			Книги
//
//	@param			book-id		path		pgtype.UUID	true	"UUID книги"
//	@param			author		body		string		false	"Полное имя автора"				Example(Известный автор)
//	@param			cover_type	body		string		false	"Тип обложки (твёрдая, мягкая)"	Example(твёрдая)
//	@param			description	body		string		false	"Описание книги"				Example(Крайне интересная книга об очень интересных вещах)
//	@param			genre		body		string		false	"Жанр книги"					Example(Роман)
//	@param			isbn		body		string		false	"ISBN книги"					Example(978-3-16-148410-0)
//	@param			language	body		string		false	"Язык книги"					Example(Русский)
//	@param			page_count	body		int			false	"Количество страниц в книге"	Example(100)
//	@param			publisher	body		string		false	"Издательство книги"			Example(Хороший Издатель)
//	@param			title		body		string		false	"Название книги"				Example(Интересная книга)
//	@param			year		body		int			false	"Год выпуска книги"				Example(2026)
//
//	@success		200			{object}	db.Book
//	@failure		400			{object}	ErrorResponse
//	@failure		409			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/books/{book-id} [patch]
func UpdateBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.UpdateBookParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	data.ID, err = util.PathUUID(consts.BookIDPathValue, request)
	if err = errors.Join(err, json.NewDecoder(request.Body).Decode(&data)); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.BookIDPathValue))
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

// SetBookCover godoc
//
//	@summary		Задать обложку
//	@description	Задаёт обложку для книги.
//	@accept			json
//	@produce		json
//
//	@tags			Книги
//
//	@param			book-id		path		pgtype.UUID	true	"UUID книги"
//	@param			imageuuid	body		string		true	"ID обложки с сервера хранилища"
//
//	@success		200			{object}	db.Book
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/books/{book-id}/cover [put]
func SetBookCover(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.SetCoverImageParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	data.ID, err = util.PathUUID(consts.BookIDPathValue, request)
	if err = errors.Join(err, json.NewDecoder(request.Body).Decode(&data)); err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.BookIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	book, err := q.SetCoverImage(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().
		Str("id", book.ID.String()).
		Str("imageUUID", book.ImageUuid.String()).
		Msg("book cover set")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(book); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// DeleteBook godoc
//
//	@summary		Удалить книгу
//	@description	Удаляет книгу из базы данных.
//	@accept			plain
//	@produce		json
//
//	@tags			Книги
//
//	@param			book-id	path		pgtype.UUID	true	"UUID книги"
//
//	@success		200		{string}	string		""
//	@failure		400		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/books/{book-id} [delete]
func DeleteBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.BookIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.BookIDPathValue))
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

// SearchBook godoc
//
//	@summary		Найти книги
//	@description	Возвращает список книг, конкретно соответствующих поисковым параметрам.
//	@accept			json
//	@produce		json
//
//	@tags			Книги
//
//	@param			limit		query		int		false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset		query		int		false	"Смещение возвращаемых элементов"			Example(5)
//
//	@param			author		query		string	false	"Полное имя автора"							Example(Известный автор)
//	@param			genre		query		string	false	"Жанр книги"								Example(Роман)
//	@param			language	query		string	false	"Язык книги"								Example(Русский)
//	@param			publisher	query		string	false	"Издательство книги"						Example(Хороший Издатель)
//	@param			title		query		string	false	"Название книги"							Example(Интересная книга)
//	@param			year		query		int		false	"Год выпуска книги"							Example(2026)
//
//	@success		200			{array}		db.Book
//	@failure		400			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/books/search [get]
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

// ListBooks godoc
//
//	@summary		Список книг
//	@description	Возвращает список всех книг, хранящихся на сервере.
//	@accept			plain
//	@produce		json
//
//	@tags			Книги
//
//	@param			limit	query		int	false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int	false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.Book
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/books [get]
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

// SearchBooksCombined godoc
//
//	@summary		Найти книги
//	@description	Возвращает список книг, соответствующих поисковому параметру. Ведётся поиск сразу по названию, автору и издательству.
//	@accept			json
//	@produce		json
//
//	@tags			Книги
//
//	@param			limit	query		int		false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int		false	"Смещение возвращаемых элементов"			Example(5)
//
//	@param			query	query		string	false	"Поисковой запрос"							Example(книг)
//
//	@success		200		{array}		db.Book
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/books/combinedSearch [get]
func SearchBooksCombined(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.SearchBooksCombinedParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	params.Query = query.Get("query")
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.SearchBooksCombined(database.Context, params)
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
