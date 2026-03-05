package sites

import (
	"housebook_scraper/types"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// ScrapeMyBook scrapes the book from MyBook.
func ScrapeMyBook(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error) {
	collector.OnHTML(".iDfqxd", func(element *colly.HTMLElement) {
		// language, publisher, cover type and color are not found on mybook.ru
		pageCountString := element.ChildText(".eSjGMZ p:nth-child(1)")
		if result.PageCount, err = strconv.Atoi(pageCountString[0:strings.Index(pageCountString, " ")]); err != nil {
			return
		}
		result.Title = element.ChildText("h1")
		result.Description = element.ChildText(".gAFRve>p")
		result.Author = element.ChildText(".eFHndH>.bkolKJ>a>.fZFkVd>.jVKkXg")
		result.ISBN = element.ChildText(".iSnZQd div:nth-child(2)>dl:nth-child(1)>dd")
		result.Genre = element.ChildText(".cFmLKi>div:nth-child(2)>.gHinNz>div:nth-child(1)>div>a")
	})
	err = collector.Visit(requestURL.String())
	return
}
