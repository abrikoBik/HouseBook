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

// CreateTrade godoc
//
//	@summary		Создать сделку
//	@description	Создаёт новую сделку.
//	@accept			json
//	@produce		json
//
//	@tags			Сделки
//
//	@param			ownerid	body		int	true	"ID инициатора сделки"
//	@param			buyerid	body		int	true	"ID другой стороны сделки"
//
//	@success		201		{object}	db.Trade
//	@failure		400		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/trades [post]
func CreateTrade(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.CreateTradeParams
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
	trade, err := q.CreateTrade(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", trade.ID).Msg("trade created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(trade); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// GetTrade godoc
//
//	@summary		Получить сделку
//	@description	Возвращает информацию о сделке.
//	@accept			plain
//	@produce		json
//
//	@tags			Сделки
//
//	@param			trade-id	path		pgtype.UUID	true	"UUID сделки"
//
//	@success		200			{object}	db.Trade
//	@failure		400			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/trades/{trade-id} [get]
func GetTrade(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.TradeIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.TradeIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	trade, err := q.GetTrade(database.Context, uuid)
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
	Logger.Info().Interface("id", trade.ID).Msg("book retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(trade); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// DeleteTrade godoc
//
//	@summary		Удалить сделку
//	@description	Удаляет сделку из базы данных.
//	@accept			plain
//	@produce		json
//
//	@tags			Сделки
//
//	@param			trade-id	path		pgtype.UUID	true	"UUID сделки"
//
//	@success		200			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/trades/{trade-id} [delete]
func DeleteTrade(writer http.ResponseWriter, request *http.Request) {
	var (
		err error
		r   = render.New()
		q   = db.New(database.Pool)
	)
	uuid, err := util.PathUUID(consts.TradeIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.TradeIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteTrade(database.Context, uuid); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", uuid).Msg("trade deleted")

	writer.WriteHeader(http.StatusOK)
}

// ListTrades godoc
//
//	@summary		Список сделок
//	@description	Возвращает список всех сделок, хранящихся на сервере.
//	@accept			plain
//	@produce		json
//
//	@tags			Сделки
//
//	@param			limit	query		int	false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int	false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.Trade
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/trades [get]
func ListTrades(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListTradesParams
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
	trades, err := q.ListTrades(database.Context, params)
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
	Logger.Info().Int("len(trades)", len(trades)).Msg("trades listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(trades); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// ListUserTrades godoc
//
//	@summary		Список сделок пользователя
//	@description	Возвращает список сделок пользователя.
//	@accept			plain
//	@produce		json
//
//	@tags			Сделки
//
//	@param			user-id	path		int	true	"ID пользователя"
//	@param			limit	query		int	false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset	query		int	false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200		{array}		db.Trade
//	@failure		400		{object}	ErrorResponse
//	@failure		404		{object}	ErrorResponse
//	@failure		500		{object}	ErrorResponse
//	@router			/api/user/{user-id}/trades [get]
func ListUserTrades(writer http.ResponseWriter, request *http.Request) {
	// TODO: move function to users
	var (
		err    error
		params db.ListUserTradesParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	params.Userid, err = util.PathID(consts.UserIDPathValue, request)
	if err != nil {
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
	trades, err := q.ListUserTrades(database.Context, params)
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
		Int64("id", params.Userid).
		Int("len(trades)", len(trades)).
		Msg("user trades listed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(trades); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}

// AddBookToTrade godoc
//
//	@summary		Добавить книгу
//	@description	Добавляет книгу в существующую сделку.
//	@accept			json
//	@produce		json
//
//	@tags			Сделки
//
//	@param			trade-id	path		pgtype.UUID	true	"UUID сделки"
//	@param			ownerid		body		int			true	"ID владельца книги"
//	@param			bookid		body		pgtype.UUID	true	"ID книги"
//
//	@success		201			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		409			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/trades/{trade-id}/books [patch]
func AddBookToTrade(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.AddBookToTradeParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	tradeID, err := util.PathUUID(consts.TradeIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.TradeIDPathValue))
		return
	}
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}
	data.Tradeid = tradeID

	if PingDB(r, writer) != nil {
		return
	}
	if err = q.AddBookToTrade(database.Context, data); err != nil {
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
		Interface("tradeID", data.Tradeid).
		Interface("bookID", data.Bookid).
		Msg("book added to trade")

	writer.WriteHeader(http.StatusCreated)
}

// ListTradeBooks godoc
//
//	@summary		Список книг в сделке
//	@description	Возвращает список всех книг, включённых в сделку.
//	@accept			plain
//	@produce		json
//
//	@tags			Сделки
//
//	@param			trade-id	path		pgtype.UUID	true	"UUID сделки"
//	@param			limit		query		int			false	"Лимит количества возвращаемых элементов"	Example(10)
//	@param			offset		query		int			false	"Смещение возвращаемых элементов"			Example(5)
//
//	@success		200			{array}		db.TradeBook
//	@failure		400			{object}	ErrorResponse
//	@failure		404			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/trades/{trade-id}/books [get]
func ListTradeBooks(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.ListTradeBooksParams
		query  = request.URL.Query()
		r      = render.New()
		q      = db.New(database.Pool)
	)
	params.Limit, params.Offset, err = util.GetSearchPagination(query)
	if err != nil {
		RenderError(r, writer, ErrorBadPaginationParams(err))
		return
	}
	tradeID, err := util.PathUUID(consts.TradeIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.TradeIDPathValue))
		return
	}
	params.ID = tradeID

	if PingDB(r, writer) != nil {
		return
	}
	books, err := q.ListTradeBooks(database.Context, params)
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

// DeleteTradeBook godoc
//
//	@summary		Удалить книгу
//	@description	Удаляет книгу из существующей сделки.
//	@accept			json
//	@produce		json
//
//	@tags			Сделки
//
//	@param			library-id	path		pgtype.UUID	true	"UUID библиотеки"
//	@param			book-id		path		pgtype.UUID	true	"UUID книги"
//
//	@success		201			{string}	string		""
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/trades/{trade-id}/books/{book-id} [delete]
func DeleteTradeBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.DeleteBookFromTradeParams
		r      = render.New()
		q      = db.New(database.Pool)
	)
	tradeID, err := util.PathUUID(consts.TradeIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.TradeIDPathValue))
		return
	}
	params.TradeID = tradeID
	bookID, err := util.PathUUID(consts.BookIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.BookIDPathValue))
		return
	}
	params.BookID = bookID

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteBookFromTrade(database.Context, params); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().
		Interface("tradeID", tradeID).
		Interface("bookID", bookID).
		Msg("book deleted from trade")

	writer.WriteHeader(http.StatusOK)
}

// ConfirmTrade godoc
//
//	@summary		Завершить сделку
//	@description	Помечает сделку завершённой со стороны пользователя.
//	@accept			json
//	@produce		json
//
//	@tags			Сделки
//
//	@param			trade-id	path		pgtype.UUID	true	"UUID сделки"
//	@param			userid		body		int			true	"ID пользователя"
//
//	@success		200			{object}	db.Trade
//	@failure		400			{object}	ErrorResponse
//	@failure		500			{object}	ErrorResponse
//	@router			/api/trades/{trade-id} [patch]
func ConfirmTrade(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.ConfirmTradeParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	if err := json.NewDecoder(request.Body).Decode(&data); err != nil {
		RenderError(r, writer, ErrorBadJSON(err))
		return
	}
	data.ID, err = util.PathUUID(consts.TradeIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.TradeIDPathValue))
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	trade, err := q.ConfirmTrade(database.Context, data)
	if err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().
		Str("id", trade.ID.String()).
		Bool("ownerConfirmed", trade.OwnerConfirmed).
		Bool("buyerConfirmed", trade.BuyerConfirmed).
		Time("completedAt", trade.CompletedAt.Time).
		Msg("trade confirmed")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(trade); err != nil {
		RenderError(r, writer, ErrorRender(err))
		return
	}
}
