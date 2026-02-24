// Package handlers contains all HTTP handlers.
package handlers

import (
	"housebook/database"
	"net/http"

	"github.com/unrolled/render"
)

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
