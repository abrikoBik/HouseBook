// Package handlers contains all HTTP handler functions.
package handlers

import (
	"context"
	"housebook_scraper/sites"
	"net/http"

	"github.com/rs/zerolog"
)

var (
	ctx    context.Context
	logger *zerolog.Logger
)

type Request struct {
	URL string
}

func Config(c context.Context) *http.ServeMux {
	ctx = c
	logger = zerolog.Ctx(ctx)
	router := http.NewServeMux()
	router.HandleFunc("GET /health", Health)
	router.HandleFunc("POST /mybook", BuildHandler(sites.MyBook, logger))
	router.HandleFunc("POST /labirint", BuildHandler(sites.Labirint, logger))
	router.HandleFunc("POST /litres", BuildHandler(sites.Litres, logger))
	return router
}

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
