package util

import (
	"errors"
	"net/http"
	"net/url"
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

// SearchParams takes a URL query and extracts limit and offset fields from it.
func SearchParams(query url.Values) (limit int, offset int, err error) {
	if limit, err = strconv.Atoi(query.Get("limit")); err != nil {
		return 0, 0, err
	}
	if offset, err = strconv.Atoi(query.Get("offset")); err != nil {
		return 0, 0, err
	}
	return
}
