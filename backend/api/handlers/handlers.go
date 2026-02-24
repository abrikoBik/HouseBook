// Package handlers contains all HTTP handlers.
package handlers

import (
	"context"
	"housebook/database"
	"net/http"

	"github.com/rs/zerolog"
	"github.com/unrolled/render"
)

var (
	Context context.Context
	Logger  *zerolog.Logger
)

func Config(ctx context.Context) {
	Context = ctx
	Logger = zerolog.Ctx(Context)
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
