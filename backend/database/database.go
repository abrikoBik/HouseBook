// Package database contains all database microservice properties and functions: start, stop, etc.
package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var (
	Context context.Context
	Pool    *pgxpool.Pool
	err     error
)

func Start(ctx context.Context, hostname string, port uint16, callback func()) {
	Context = ctx
	// TODO: remove sslmode=disable
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		hostname,
		port,
		os.Getenv("POSTGRES_DB"),
	)
	if Pool, err = pgxpool.New(Context, connString); err != nil {
		panic(err)
	}
	if err = Pool.Ping(Context); err != nil {
		panic(err)
	}
	callback()
}

func Stop(callback func()) {
	Pool.Close()
	callback()
}
