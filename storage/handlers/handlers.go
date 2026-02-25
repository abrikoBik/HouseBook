// Package handlers contains all HTTP handler functions.
package handlers

import (
	"context"
	"net/http"

	"github.com/rs/zerolog"
)

var (
	ctx    context.Context
	logger *zerolog.Logger
)

func Config(c context.Context) *http.ServeMux {
	ctx = c
	logger = zerolog.Ctx(ctx)
	router := http.NewServeMux()
	router.HandleFunc("GET /health", Health)
	router.HandleFunc("POST /", ReceiveFile)
	router.HandleFunc("GET /{id}", SendFile)
	return router
}

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
