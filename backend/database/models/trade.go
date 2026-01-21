package models

import (
	"housebook/database"
	"net/http"
	"time"

	"github.com/emvi/null"
	"github.com/google/uuid"
	"github.com/uptrace/bun"
)

type Trade struct {
	bun.BaseModel `bun:"table:trades"`
	UUID          uuid.UUID `bun:",pk,type:uuid,default:uuid_generate_v4()" json:"uuid"`

	Seller User   `bun:"rel:has-one,join:id=id" json:"seller"`
	Buyer  User   `bun:"rel:has-one,join:id=id" json:"buyer"`
	Books  []Book `bun:"rel:has-many,join:id=id" json:"books"`

	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	CompletedAt null.Time `json:"completedAt"`
}

type TradeRequest struct {
	*Trade
	SellerID int64      `json:"sellerID"`
	BuyerID  int64      `json:"buyerID"`
	BookIDs  uuid.UUIDs `json:"bookIDs"`
}

type TradeResponse struct {
	*Trade
}

func (trade *Trade) Render(w http.ResponseWriter, r *http.Request) error {
	// TODO: add proper error handling
	return nil
}

func (trade *TradeRequest) Bind(r *http.Request) error {
	// TODO: add proper error handling
	return nil
}

func (trade *Trade) Insert(request TradeRequest) error {
	trade = request.Trade
	_, err := database.BunDatabase.
		NewInsert().
		Model(trade).
		Exec(database.Context)
	return err
}

func (trade *Trade) Select() error {
	return database.BunDatabase.
		NewSelect().
		Model(trade).
		WherePK().
		Relation("Seller").
		Relation("Buyer").
		Relation("Books").
		Scan(database.Context)
}

func (trade *Trade) Delete() error {
	_, err := database.BunDatabase.
		NewDelete().
		Model(trade).
		WherePK().
		Exec(database.Context)
	return err
}
