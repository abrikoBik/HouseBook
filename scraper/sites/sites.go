// Package sites contains scrapers for supported websites.
package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"housebook_scraper/util"
	"net/url"
	"strconv"
	"sync"

	"github.com/gocolly/colly/v2"
	"github.com/gsiems/go-isbn/pkg/isbn"
	"github.com/rs/zerolog"
)

// Site contains the domain and all Colly scraper strings for a specific site.
type Site struct {
	Title, Author, ISBN, Genre, Publisher, // required fields
	Year, Language, CoverType, PageCount, Description, CoverImage, CoverImageProperty, // non-required fields
	Domain, Root string
	Quirks map[string]func(book *types.TempBook)
}

// Scrape implements a common scraper for all sites to use.
func (site *Site) Scrape(collector *colly.Collector, requestURL *url.URL, logger *zerolog.Logger) (result types.Book, err error) {
	// TODO: eliminate doubled code
	collector.OnHTML(site.Root, func(element *colly.HTMLElement) {
		result.StoreLink = requestURL.String()
		var temp types.TempBook
		temp.Title = element.ChildText(site.Title)
		temp.Author = element.ChildText(site.Author)
		temp.ISBN = element.ChildText(site.ISBN)
		temp.Genre = element.ChildText(site.Genre)
		temp.Publisher = element.ChildText(site.Publisher)
		temp.Year = element.ChildText(site.Year)
		temp.Language = element.ChildText(site.Language)
		temp.CoverType = element.ChildText(site.CoverType)
		temp.PageCount = element.ChildText(site.PageCount)
		temp.Description = element.ChildText(site.Description)
		temp.ImageLink = element.ChildAttr(site.CoverImage, site.CoverImageProperty)

		// apply site-specific quirks
		var wg sync.WaitGroup
		wg.Add(len(site.Quirks))
		for name, quirk := range site.Quirks {
			go func() {
				logger.Info().Str("name", name).Msg("applying quirk")
				quirk(&temp)
				logger.Info().Str("name", name).Msg("applied quirk")
				wg.Done()
			}()
		}
		wg.Wait()
		// prepare resulting book
		result.Title = temp.Title
		result.Author = temp.Author
		// remove all foreign characters from an ISBN
		i, nerr := util.ExtractNumber(temp.ISBN)
		if nerr != nil {
			err = consts.ErrInvalidISBN
			return
		}
		result.ISBN = strconv.Itoa(i)
		if !isbn.ValidateCheckDigit(result.ISBN) {
			err = consts.ErrInvalidISBN
			return
		}
		result.Genre = temp.Genre
		result.Publisher = temp.Publisher
		if result.Year, err = util.ExtractNumber(temp.Year); err != nil {
			return
		}
		result.Language = temp.Language
		result.CoverType = temp.CoverType
		result.Description = temp.Description
		if result.PageCount, err = util.ExtractNumber(temp.PageCount); err != nil {
			return
		}
		u, err := url.Parse(temp.ImageLink)
		if err != nil {
			return
		}
		if u.String() != "" {
			if u.Scheme == "" {
				u.Scheme = "http"
			}
			result.ImageLink = u.String()
		}

		// go over required fields
		required := []*string{
			&result.Title,
			&result.Author,
			&result.Genre,
			&result.Publisher,
		}
		for _, field := range required {
			if *field == "" {
				*field = "Неизвестен"
			}
		}
	})
	err = collector.Visit(requestURL.String())
	collector.Wait()
	return
}
