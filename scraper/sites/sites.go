// Package sites contains scrapers for supported websites.
package sites

import (
	"housebook_scraper/types"
	"net/url"

	"github.com/gocolly/colly/v2"
)

type Scraper func(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error)
