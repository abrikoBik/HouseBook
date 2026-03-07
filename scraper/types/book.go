// Package types contains all types used by the scraper service.
package types

type Book struct {
	Title       string
	Author      string
	ISBN        string
	Genre       string
	Language    string
	Publisher   string
	CoverType   string
	PageCount   int
	Description string
	Color       bool
}
