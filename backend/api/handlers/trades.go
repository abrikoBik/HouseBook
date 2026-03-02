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

func GetTrade(writer http.ResponseWriter, request *http.Request) {
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

func DeleteTrade(writer http.ResponseWriter, request *http.Request) {
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
	if err := q.DeleteTrade(database.Context, uuid); err != nil {
		RenderError(r, writer, ErrorGeneric(err))
		return
	}
	Logger.Info().Interface("id", uuid).Msg("trade deleted")

	writer.WriteHeader(http.StatusOK)
}

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

func AddBookToTrade(writer http.ResponseWriter, request *http.Request) {
	var (
		err  error
		data db.AddBookToTradeParams
		r    = render.New()
		q    = db.New(database.Pool)
	)
	tradeID, err := util.PathUUID(consts.UUIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.UUIDPathValue))
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
	tradeID, err := util.PathUUID(consts.UUIDPathValue, request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, consts.UUIDPathValue))
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

func DeleteTradeBook(writer http.ResponseWriter, request *http.Request) {
	var (
		err    error
		params db.DeleteBookFromTradeParams
		r      = render.New()
		q      = db.New(database.Pool)
	)
	tradeID, err := util.PathUUID("trade-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "trade-uuid"))
		return
	}
	params.TradeID = tradeID
	bookID, err := util.PathUUID("book-uuid", request)
	if err != nil {
		RenderError(r, writer, ErrorBadID(err, "book-uuid"))
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
