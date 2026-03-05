package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"net/url"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

// ScrapeLabirint scrapes the book from Labirint.
func ScrapeLabirint(collector *colly.Collector, requestURL *url.URL) (result types.Book, err error) {
	collector.OnHTML("body", func(element *colly.HTMLElement) {
		// genre, color are not found on litres.ru
		pageCountString := element.ChildText("._right_u86in_12>div:nth-child(4)>.text-black>span:nth-child(1)")
		pageCountSlice := strings.Split(pageCountString, " ")
		if len(pageCountSlice) == 0 {
			err = consts.ErrPageCountEmpty
			return
		}
		if result.PageCount, err = strconv.Atoi(pageCountSlice[0]); err != nil {
			return
		}
		result.Title = element.ChildText("h1")
		result.Author = element.ChildText("._features_ctofl_78>div:nth-child(1)>div:nth-child(2)>a")
		result.ISBN = element.ChildText("._right_u86in_12>div:nth-child(2)>.text-black>span")
		result.Language = element.ChildText("._right_u86in_12>div:nth-child(3)>.text-black>span")
		result.Publisher = element.ChildText("._left_u86in_12>div:nth-child(4)>div:nth-child(2)>a")
		result.CoverType = element.ChildText("._right_u86in_12>div:nth-child(7)>.text-black>span")
		result.Description = element.ChildText("._wrapper_1rsml_1>div:nth-child(1)>div")
	})
	err = collector.Visit(requestURL.String())
	return
}
