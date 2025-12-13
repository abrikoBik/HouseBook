// Package database contains all database microservice properties and functions: start, stop, etc.
package database

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun/driver/pgdriver"
)

var (
	Context  context.Context
	Database *sql.DB
)

func Start(hostname string, port uint16) error {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		hostname,
		port,
		os.Getenv("POSTGRES_DB"),
	)
	Conn := sql.OpenDB(pgdriver.NewConnector(
		pgdriver.WithDSN(url),
	))
	defer Conn.Close()
	return nil
}
