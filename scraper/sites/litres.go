package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"strings"
)

var Litres = Site{
	Domain:             consts.LitresDomain,
	Root:               "body",
	Title:              ".BookCard_book__title__EVxWV",
	Author:             "li.Breadcrumbs_breadcrumbs__item__zPK82:nth-child(3) > a:nth-child(1) > span:nth-child(2)",
	ISBN:               "div.CharacteristicsBlock_characteristic__4pi7v:nth-child(5) > span:nth-child(2)[itemprop=\"isbn\"]",
	Genre:              "li.Breadcrumbs_breadcrumbs__item__zPK82:nth-child(2) > a:nth-child(1) > span:nth-child(1)",
	Publisher:          "div.CharacteristicsBlock_characteristic__4pi7v:nth-child(7) > a:nth-child(2)",
	Year:               "div.CharacteristicsBlock_characteristic__4pi7v:nth-child(3) > span:nth-child(2)",
	PageCount:          "div.CharacteristicsBlock_characteristic__4pi7v:nth-child(4) > span:nth-child(2)",
	Description:        ".BookCard_book__annotation__8wq0r > p:nth-child(2)",
	CoverImage:         ".BookCover_img__ncTI5",
	CoverImageProperty: "src",
	Quirks: map[string]func(book *types.TempBook){
		"extractPageCount": func(book *types.TempBook) {
			book.PageCount = strings.Split(book.PageCount, " ")[0]
		},
		"extractYear": func(book *types.TempBook) {
			split := strings.Split(book.PageCount, " ")
			book.Year = split[len(split)-1]
		},
	},
}
