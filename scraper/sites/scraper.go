package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Scraper interface {
	Scrape(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error)
}

// Site contains all Colly scraper strings for a specific site.
type Site struct {
	Root,
	Title,
	Author,
	ISBN,
	Genre,
	Language,
	Publisher,
	CoverType,
	PageCount,
	Description,
	Color string
}

// Scrape implements a common scraper for all sites to use.
func (site *Site) Scrape(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error) {
	collector.OnHTML(site.Root, func(element *colly.HTMLElement) {
		result.Title = element.ChildText(site.Title)
		result.Author = element.ChildText(site.Author)
		result.ISBN = element.ChildText(site.ISBN)
		result.Genre = element.ChildText(site.Genre)
		result.Language = element.ChildText(site.Language)
		result.Publisher = element.ChildText(site.Publisher)
		result.CoverType = element.ChildText(site.CoverType)
		pageCountString := element.ChildText(site.PageCount)
		pageCountSlice := strings.Split(pageCountString, " ")
		if len(pageCountSlice) == 0 {
			err = consts.ErrPageCountEmpty
			return
		}
		if result.PageCount, err = strconv.Atoi(pageCountSlice[0]); err != nil {
			return
		}
		_ = element.ChildText(site.Color)
		result.Description = element.ChildText(site.Description)
	})
	err = collector.Visit(requestURL.String())
	return
}
