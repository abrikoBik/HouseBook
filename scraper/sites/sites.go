// Package sites contains scrapers for supported websites.
package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"housebook_scraper/util"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
	"github.com/gsiems/go-isbn/pkg/isbn"
)

// Site contains the domain and all Colly scraper strings for a specific site.
type Site struct {
	Title, Author, ISBN, Genre, Publisher, // required fields
	Year, Language, CoverType, PageCount, Description, CoverImage, CoverImageProperty, // non-required fields
	Domain, Root string
	Quirks map[string]func()
}

// Scrape implements a common scraper for all sites to use.
func (site *Site) Scrape(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error) {
	// TODO: eliminate doubled code
	collector.OnHTML(site.Root, func(element *colly.HTMLElement) {
		result.StoreLink = requestURL.String()
		result.Title = element.ChildText(site.Title)
		result.Author = element.ChildText(site.Author)
		// remove all foreign characters from an ISBN
		i, nerr := util.ExtractNumber(element.ChildText(site.ISBN))
		if nerr != nil {
			err = consts.ErrInvalidISBN
			return
		}
		result.ISBN = strconv.Itoa(i)
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
			if result.Year, err = util.ExtractNumber(year); err != nil {
				return
			}
		}
		result.Language = element.ChildText(site.Language)
		result.CoverType = element.ChildText(site.CoverType)
		if pc := element.ChildText(site.PageCount); pc != "" {
			if result.PageCount, err = util.ExtractNumber(pc); err != nil {
				return
			}
		}
		result.Description = element.ChildText(site.Description)
		imageLink := element.ChildAttr(site.CoverImage, site.CoverImageProperty)
		u, err := url.Parse(strings.Split(imageLink, " ")[0])
		if err != nil {
			return
		}
		if u.String() != "" {
			if u.Scheme == "" {
				u.Scheme = "http"
			}
			result.ImageLink = u.String()
		}
		// TODO: find a better solution
		for _, quirk := range site.Quirks {
			quirk()
		}
	})
	err = collector.Visit(requestURL.String())
	collector.Wait()
	return
}
