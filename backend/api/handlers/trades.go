package handlers

import (
	"net/http"
)

// TODO: UpdateTrade for closing trades

func CreateTrade(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	trade = &models.Trade{}
	// 	data  = &models.TradeRequest{Trade: trade}
	// )
	// if err := render.Bind(r, data); err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }

	// // look up all entities by id
	// data.Seller = models.User{ID: data.SellerID}
	// data.Buyer = models.User{ID: data.BuyerID}
	// if err := data.Seller.Select(); err != nil {
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// if err := data.Buyer.Select(); err != nil {
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// for _, uuid := range data.BookIDs {
	// 	book := models.Book{UUID: uuid}
	// 	if err := book.Select(); err != nil {
	// 		render.Render(w, r, ErrorGeneric(err))
	// 		return
	// 	}
	// 	data.Books = append(data.Books, book)
	// }

	// if err := trade.Insert(*data); err != nil {
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusCreated)
	// if err := render.Render(w, r, tradeResponse(trade)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}

func GetTrade(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	trade   = &models.Trade{}
	// 	id, err = util.PathUUID(consts.UUIDPathValue, r)
	// )
	// if err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// trade.UUID = *id
	// if err := trade.Select(); err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		render.Render(w, r, ErrorNotFound(err))
	// 		return
	// 	}
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusOK)
	// if err := render.Render(w, r, tradeResponse(trade)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}

func DeleteTrade(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	trade   = &models.Trade{}
	// 	id, err = util.PathUUID(consts.UUIDPathValue, r)
	// )
	// if err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// trade.UUID = *id
	// if err := trade.Delete(); err != nil {
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// w.WriteHeader(http.StatusOK)
	w.WriteHeader(http.StatusNotImplemented)
}

func SearchTrade(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	params = util.SearchTable{}
	// 	query  = r.URL.Query()
	// )
	// for _, param := range consts.TradeSearchParams {
	// 	params[param] = query.Get(param)
	// }
	// trades, err := models.Search[models.Trade](params)
	// if err != nil {
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		render.Render(w, r, ErrorNotFound(err))
	// 		return
	// 	}
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusOK)
	// if err := render.RenderList(w, r, tradeListResponse(trades)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}

func ListTrades(w http.ResponseWriter, r *http.Request) {
	// var (
	// 	query              = r.URL.Query()
	// 	limit, offset, err = util.SearchParams(query)
	// )
	// if err != nil {
	// 	render.Render(w, r, ErrorBadRequest(err))
	// 	return
	// }
	// trades, err := models.List[models.Trade](limit, offset)
	// if err != nil {
	// 	// TODO: decide if 404 is applicable here
	// 	if errors.Is(err, sql.ErrNoRows) {
	// 		render.Render(w, r, ErrorNotFound(err))
	// 		return
	// 	}
	// 	render.Render(w, r, ErrorGeneric(err))
	// 	return
	// }
	// render.Status(r, http.StatusOK)
	// if err := render.RenderList(w, r, tradeListResponse(trades)); err != nil {
	// 	render.Render(w, r, ErrorRender(err))
	// 	return
	// }
	w.WriteHeader(http.StatusNotImplemented)
}
