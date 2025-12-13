// Package models contains all of the models used in the database, as well as all CRUD operations related to them.
package models

import (
	"context"
	"housebook/database"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
)

// EPIC global state
var (
	ctx context.Context
	db  *bun.DB
)

func init() {
	ctx = database.Context
	db = bun.NewDB(database.Database, pgdialect.New())
}
