// Package database contains all database microservice properties and functions: start, stop, etc.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	Context     context.Context
	Database    *sql.DB
	BunDatabase *bun.DB
)

func Start(hostname string, port uint16) error {
	// TODO: remove sslmode=disable
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		hostname,
		port,
		os.Getenv("POSTGRES_DB"),
	)
	Database = sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(url),
	))
	BunDatabase = bun.NewDB(Database, pgdialect.New())
	return nil
}
