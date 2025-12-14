package util

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

// PathID adds basic error checking, as well as integer conversion for the respective URL parameter.
func PathID(name string, request *http.Request) (int64, error) {
	s := chi.URLParam(request, name)
	if s == "" {
		return 0, errors.New("id not provided")
	}
	id, err := strconv.ParseInt(s, 10, 64)
	return id, err
}

// PathUUID adds basic error checking, as well as uuid.UUID conversion for the respective URL parameter.
func PathUUID(name string, request *http.Request) (*uuid.UUID, error) {
	s := chi.URLParam(request, name)
	if s == "" {
		return nil, errors.New("id not provided")
	}
	id, err := uuid.Parse(s)
	return &id, err
}
