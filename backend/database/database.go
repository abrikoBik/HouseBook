// Package database contains all database microservice properties and functions: start, stop, etc.
package database

import (
	"context"
	"fmt"
	"housebook/database/generated"
	"os"

	"github.com/jackc/pgx/v5"
)

var (
	Context    context.Context
	Connection *pgx.Conn
	Queries    *generated.Queries
)

func Start(hostname string, port uint16) error {
	// TODO: remove sslmode=disable
	connString := fmt.Sprintf("user=%s password=%s hostname=%s port=%d db-name=%s sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		hostname,
		port,
		os.Getenv("POSTGRES_DB"),
	)
	Connection, err := pgx.Connect(Context, connString)
	if err != nil {
		return err
	}
	defer Connection.Close(Context)
	Queries = generated.New(Connection)
	return nil
}
