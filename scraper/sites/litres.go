package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// ScrapeLitres scrapes the book from Litres.
func ScrapeLitres(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error) {
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// language, cover type and color are not found on litres.ru
		pageCountString := element.ChildText("._9d4e17ed>div:nth-child(4)>span")
		pageCountSlice := strings.Split(pageCountString, " ")
		if len(pageCountSlice) == 0 {
			err = consts.ErrPageCountEmpty
			return
		}
		if result.PageCount, err = strconv.Atoi(pageCountSlice[0]); err != nil {
			return
		}
		result.Title = element.ChildText("h1")
		result.Author = element.ChildText("._86af713b>div>span")
		result.ISBN = element.ChildText("._9d4e17ed>div:nth-child(5)>span")
		result.Genre = element.ChildText(".aa63d864>div:nth-child(1)>.f12c0b54")
		result.Publisher = element.ChildText("._9d4e17ed>div:nth-child(6)>a")
		result.Description = element.ChildText(".ac83cc29>div:nth-child(1)")
	})
	err = collector.Visit(requestURL.String())
	return
}
