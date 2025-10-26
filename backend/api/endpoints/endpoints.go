// Package endpoints contains all HTTP endpoints and their routes as strings.
package endpoints

import (
	"housebook/api/handlers"
	"net/http"
)

const (
	Health = "/health"
)

func Register(mux *http.ServeMux) {
	mux.HandleFunc(Health, handlers.Health)
}
