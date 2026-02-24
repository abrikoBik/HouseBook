// Package database contains all database microservice properties and functions: start, stop, etc.
package database

import (
	"context"
	"fmt"
	"housebook/consts"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var (
	Context              context.Context
	Pool                 *pgxpool.Pool
	ErrLogger, OutLogger zerolog.Logger
	err                  error
)

func Start(ctx context.Context, hostname string, port uint16, callback func()) {
	// TODO: move this from Start()?
	Context = ctx
	ErrLogger = Context.Value(consts.ErrLogger).(zerolog.Logger)
	OutLogger = Context.Value(consts.OutLogger).(zerolog.Logger)

	// TODO: remove sslmode=disable
	connString := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		os.Getenv("POSTGRES_USER"),
		os.Getenv("POSTGRES_PASSWORD"),
		hostname,
		port,
		os.Getenv("POSTGRES_DB"),
	)
	if Pool, err = pgxpool.New(Context, connString); err != nil {
		ErrLogger.Panic().Err(err).Msg("panicking in DB (pool)")
	}
	if err = Pool.Ping(Context); err != nil {
		ErrLogger.Panic().Err(err).Msg("panicking in DB (ping)")
	}
	OutLogger.Info().Str("hostname", hostname).Uint16("port", port).Msg("started DB")
	callback()
}

func Stop(callback func()) {
	OutLogger.Info().Msg("closing DB pool")
	Pool.Close()
	callback()
}
