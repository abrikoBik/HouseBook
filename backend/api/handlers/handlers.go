// Package handlers contains all HTTP handlers.
package handlers

import (
	"context"
	"housebook/consts"
	"housebook/database"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/unrolled/render"
)

var (
	Context              context.Context
	ErrLogger, OutLogger zerolog.Logger
)

func Config(ctx context.Context) {
	Context = ctx
	ErrLogger = Context.Value(consts.ErrLogger).(zerolog.Logger)
	OutLogger = Context.Value(consts.OutLogger).(zerolog.Logger)
}

func Health(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func PingDB(renderer *render.Render, writer http.ResponseWriter) error {
	if err := database.Pool.Ping(database.Context); err != nil {
		toWrite := ErrorDatabase(err)
		renderer.JSON(writer, toWrite.Status, toWrite)
		return err
	}
	return nil
}
