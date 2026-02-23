package consts

import (
	"errors"
	"housebook/database/db"
	"net/url"
)

var (
	BookSearchParams = []string{"title", "author", "genre", "publisher", "year", "language"}
	// TODO: allow book searches
	TradeSearchParams = []string{"seller", "buyer"}
)

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
