// Package sites contains scrapers for supported websites.
package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gsiems/go-isbn/pkg/isbn"
)

// Site contains the domain and all Colly scraper strings for a specific site.
type Site struct {
	Title, Author, ISBN, Genre, Publisher, // required fields
	Year, Language, CoverType, PageCount, Description, // non-required fields
	Domain, Root string
}

// Scrape implements a common scraper for all sites to use.
func (site *Site) Scrape(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error) {
	collector.OnHTML(site.Root, func(element *colly.HTMLElement) {
		result.Title = element.ChildText(site.Title)
		result.Author = element.ChildText(site.Author)
		result.ISBN = element.ChildText(site.ISBN)
		if !isbn.ValidateCheckDigit(result.ISBN) {
			err = consts.ErrInvalidISBN
			return
		}
		result.Genre = element.ChildText(site.Genre)
		result.Publisher = element.ChildText(site.Publisher)

		if year := element.ChildText(site.Year); year != "" {
			if result.Year, err = strconv.Atoi(strings.Split(year, " ")[0]); err != nil {
				return
			}
		}
		result.Language = element.ChildText(site.Language)
		result.CoverType = element.ChildText(site.CoverType)
		if pc := element.ChildText(site.PageCount); pc != "" {
			if result.PageCount, err = strconv.Atoi(strings.Split(pc, " ")[0]); err != nil {
				return
			}
		}
		result.Description = element.ChildText(site.Description)
	})
	err = collector.Visit(requestURL.String())
	collector.Wait()
	return
}
