package util

import (
	"errors"
	"housebook/database/db"
	"net/http"
	"net/url"
	"strconv"

	"github.com/go-chi/chi/v5"
	"github.com/jackc/pgx/v5/pgtype"
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

// PathUUID adds basic error checking, as well as pgtype.UUID conversion for the respective URL parameter.
func PathUUID(name string, request *http.Request) (result pgtype.UUID, err error) {
	s := chi.URLParam(request, name)
	if s == "" {
		err = errors.New("uuid not provided")
		return
	}
	err = result.Scan(s)
	return
}

// GetSearchPagination takes a URL query and extracts limit and offset fields from it for pagination.
// If limit and offset are not provided in the URL query, they are set to 10 and 0 respectively.
func GetSearchPagination(query url.Values) (int32, int32, error) {
	var (
		limitQuery        = query.Get("limit")
		offsetQuery       = query.Get("offset")
		limit       int32 = 10
		offset      int32
	)
	if limitQuery != "" {
		converted, err := strconv.Atoi(limitQuery)
		if err != nil {
			return 0, 0, nil
		}
		limit = int32(converted)
	}
	if offsetQuery != "" {
		converted, err := strconv.Atoi(offsetQuery)
		if err != nil {
			return 0, 0, nil
		}
		offset = int32(converted)
	}
	return limit, offset, nil
}

// GetSearchParams scans search parameters from query into a db.SearchBooksParams struct.
func GetSearchParams(query url.Values) (result db.SearchBooksParams, err error) {
	// seems to be the only way to do it
	err = errors.Join(err, result.Author.Scan(query.Get("author")))
	err = errors.Join(err, result.Genre.Scan(query.Get("genre")))
	err = errors.Join(err, result.Language.Scan(query.Get("language")))
	err = errors.Join(err, result.Publisher.Scan(query.Get("publisher")))
	err = errors.Join(err, result.Title.Scan(query.Get("title")))
	// ugly hack
	y := query.Get("year")
	if y == "" {
		y = "0"
	}
	err = errors.Join(err, result.Year.Scan(y))
	return
}
