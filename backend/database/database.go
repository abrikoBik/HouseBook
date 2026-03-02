// Package database contains all database microservice properties and functions: start, stop, etc.
package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var (
	Context context.Context
	Logger  *zerolog.Logger
	Pool    *pgxpool.Pool
	err     error
)

func Config(ctx context.Context, callback func()) {
	Context = ctx
	Logger = zerolog.Ctx(Context)
	callback()
}

func Start(hostname string, port uint16, callback func()) {
	// TODO: remove sslmode=disable
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		hostname,
		port,
		os.Getenv("POSTGRES_DB"),
	)
	Logger.Info().Str("hostname", hostname).Uint16("port", port).Msg("starting DB")
	if Pool, err = pgxpool.New(Context, connString); err != nil {
		Logger.Panic().Err(err).Msg("panicking in DB (pool)")
	}
	if err = Pool.Ping(Context); err != nil {
		Logger.Panic().Err(err).Msg("panicking in DB (ping)")
	}
	callback()
}

func Stop(callback func()) {
	Logger.Info().Msg("closing DB pool")
	Pool.Close()
	callback()
}
