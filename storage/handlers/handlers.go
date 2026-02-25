// Package handlers contains all HTTP handler functions.
package handlers

import (
	"context"
	"net/http"
)

var ctx context.Context

func Config(c context.Context) *http.ServeMux {
	ctx = c
	router := http.NewServeMux()
	router.HandleFunc("GET /health", Health)
	router.HandleFunc("POST /", ReceiveFile)
	router.HandleFunc("GET /{id}", SendFile)
	return router
}

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
