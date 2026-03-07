// Package consts contains all non-changing values required by the scraper service, such as hostname, allowed domains for scraping, etc.
package consts

import "errors"

const (
	Hostname = "0.0.0.0"
	Port     = 3002

	MyBookDomain   = "mybook.ru"
	LitresDomain   = "www.litres.ru"
	LabirintDomain = "www.labirint.ru"
)

var (
	ErrPageCountEmpty = errors.New("page count is empty")
)
