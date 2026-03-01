package handlers

import (
	"database/sql"
	"encoding/json"
	"housebook/consts"
	"housebook/database"
	"housebook/database/db"
	"housebook/util"
	"net/http"

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
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	trade, err := q.CreateTrade(database.Context, data)
	if err != nil {
		var toWrite *ErrorResponse
		switch err {
		default:
			toWrite = ErrorGeneric(err)
		}
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	Logger.Info().Interface("id", trade.ID).Msg("trade created")

	writer.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(writer).Encode(trade); err != nil {
		toWrite := ErrorRender(err)
		r.JSON(writer, toWrite.Status, toWrite)
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
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	trade, err := q.GetTrade(database.Context, uuid)
	if err != nil {
		var toWrite *ErrorResponse
		switch err {
		case sql.ErrNoRows:
			toWrite = ErrorNotFound(err)
		default:
			toWrite = ErrorGeneric(err)
		}
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	Logger.Info().Interface("id", trade.Trade.ID).Msg("book retrieved")

	writer.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(writer).Encode(trade); err != nil {
		toWrite := ErrorRender(err)
		r.JSON(writer, toWrite.Status, toWrite)
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
		toWrite := ErrorBadRequest(err)
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}

	if PingDB(r, writer) != nil {
		return
	}
	if err := q.DeleteTrade(database.Context, uuid); err != nil {
		var toWrite *ErrorResponse
		switch err {
		default:
			toWrite = ErrorGeneric(err)
		}
		r.JSON(writer, toWrite.Status, toWrite)
		return
	}
	Logger.Info().Interface("id", uuid).Msg("trade deleted")

	writer.WriteHeader(http.StatusOK)
}

func ListTrades(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}

func AddBookToTrade(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusNotImplemented)
}
