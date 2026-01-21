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

	Seller User   `bun:",notnull" json:"seller"`
	Buyer  User   `bun:",notnull" json:"buyer"`
	Books  []Book `bun:",notnull" json:"books"`

	CreatedAt   time.Time `bun:",nullzero,notnull,default:current_timestamp" json:"createdAt"`
	CompletedAt null.Time `json:"completedAt"`
}

type TradeRequest struct {
	*Trade
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
