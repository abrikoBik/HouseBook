// Package db contains all database microservice properties and functions: start, stop, etc.
package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Context context.Context
	Conn    pgxpool.Conn
)

func Start(hostname string, port uint16) error {
	url := fmt.Sprintf("postgres://%s:%s@%s:%d/%s",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		hostname,
		port,
		os.Getenv("POSTGRES_DB"),
	)
	Conn, err := pgxpool.New(Context, url)
	if err != nil {
		return err
	}
	defer Conn.Close()
	return nil
}
