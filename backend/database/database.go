// Package database contains all database microservice properties and functions: start, stop, etc.
package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

var (
	Context context.Context
	Logger  *zerolog.Logger
	Pool    *pgxpool.Pool

	connString string
	err        error
)

type Conf struct {
	Context                                context.Context
	Username, Password, Hostname, Database string
	Port                                   uint16
}

func Config(config Conf, callback func()) {
	Context = config.Context
	Logger = zerolog.Ctx(Context)
	// TODO: remove sslmode=disable
	connString = fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		config.Username,
		config.Password,
		config.Hostname,
		config.Port,
		config.Database,
	)
	callback()
}

func Start(hostname string, port uint16, callback func()) {
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
