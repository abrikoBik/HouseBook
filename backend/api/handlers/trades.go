package handlers

import (
	"database/sql"
	"errors"
	"housebook/consts"
	"housebook/database/models"
	"housebook/util"
	"net/http"
	"strconv"

	"github.com/go-chi/render"
)

func tradeResponse(trade *models.Trade) *models.TradeResponse {
	return &models.TradeResponse{Trade: trade}
}

func tradeListResponse(trades []*models.Trade) (result []render.Renderer) {
	for _, trade := range trades {
		result = append(result, tradeResponse(trade))
	}
	return
}

func CreateTrade(w http.ResponseWriter, r *http.Request) {
	var (
		trade = &models.Trade{}
		data  = &models.TradeRequest{}
	)
	if err := render.Bind(r, data); err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}

	// look up all entities by id
	data.Seller.ID = data.SellerID
	data.Buyer.ID = data.BuyerID
	if err := data.Seller.Select(); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	if err := data.Buyer.Select(); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	for _, uuid := range data.BookIDs {
		book := models.Book{UUID: uuid}
		if err := book.Select(); err != nil {
			render.Render(w, r, ErrorGeneric(err))
			return
		}
		data.Books = append(data.Books, book)
	}

	if err := trade.Insert(*data); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusCreated)
	if err := render.Render(w, r, tradeResponse(trade)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func GetTrade(w http.ResponseWriter, r *http.Request) {
	var (
		trade   = &models.Trade{}
		id, err = util.PathUUID(consts.UUIDPathValue, r)
	)
	if err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	trade.UUID = *id
	if err := trade.Select(); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			render.Render(w, r, ErrorNotFound(err))
			return
		}
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.Render(w, r, tradeResponse(trade)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func DeleteTrade(w http.ResponseWriter, r *http.Request) {
	var (
		trade   = &models.Trade{}
		id, err = util.PathUUID(consts.UUIDPathValue, r)
	)
	if err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	trade.UUID = *id
	if err := trade.Delete(); err != nil {
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	w.WriteHeader(http.StatusOK)
}

func SearchTrade(w http.ResponseWriter, r *http.Request) {
	var (
		params = util.SearchTable{}
		query  = r.URL.Query()
	)
	for _, param := range consts.TradeSearchParams {
		params[param] = query.Get(param)
	}
	trades, err := models.Search[models.Trade](params)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			render.Render(w, r, ErrorNotFound(err))
			return
		}
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.RenderList(w, r, tradeListResponse(trades)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}

func ListTrades(w http.ResponseWriter, r *http.Request) {
	var (
		limit, offset int
		err           error
		query         = r.URL.Query()
	)
	// TODO: convert to a function
	if limit, err = strconv.Atoi(query.Get("limit")); err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}
	if offset, err = strconv.Atoi(query.Get("offset")); err != nil {
		render.Render(w, r, ErrorBadRequest(err))
		return
	}

	trades, err := models.List[models.Trade](limit, offset)
	if err != nil {
		// TODO: decide if 404 is applicable here
		if errors.Is(err, sql.ErrNoRows) {
			render.Render(w, r, ErrorNotFound(err))
			return
		}
		render.Render(w, r, ErrorGeneric(err))
		return
	}
	render.Status(r, http.StatusOK)
	if err := render.RenderList(w, r, tradeListResponse(trades)); err != nil {
		render.Render(w, r, ErrorRender(err))
		return
	}
}
