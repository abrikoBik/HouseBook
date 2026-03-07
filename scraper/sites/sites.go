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
	// TODO: eliminate doubled code
	collector.OnHTML(site.Root, func(element *colly.HTMLElement) {
		result.Title = element.ChildText(site.Title)
		result.Author = element.ChildText(site.Author)
		result.ISBN = element.ChildText(site.ISBN)
		if !isbn.ValidateCheckDigit(result.ISBN) {
			err = consts.ErrInvalidISBN
			return
		}
		// this is a very dirty hack, some sites do not provide a genre and a publisher
		// TODO: check on the backend if those fields really should be non null
		genre := element.ChildText(site.Genre)
		if genre == "" {
			genre = "Неизвестен"
		}
		result.Genre = genre
		publisher := element.ChildText(site.Publisher)
		if publisher == "" {
			publisher = "Неизвестен"
		}
		result.Publisher = publisher

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
