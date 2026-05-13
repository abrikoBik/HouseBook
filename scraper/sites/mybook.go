package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"strings"
)

var MyBook = Site{
	Domain:      consts.MyBookDomain,
	Root:        "body",
	Title:       "h1.sc-bdfBwQ",
	Author:      ".sc-1c0xbiw-8>div:nth-child(1)>a:nth-child(1)>div:nth-child(1)>div:nth-child(1)",
	ISBN:        "div.iszfik-15:nth-child(2)>dl:nth-child(1)>dd:nth-child(2)",
	Genre:       "div.sc-1sg8rha-0:nth-child(2)>div:nth-child(1)>div:nth-child(1)>a:nth-child(1)",
	Year:        "div.iszfik-15:nth-child(1)>dl:nth-child(1)>dd:nth-child(2)",
	PageCount:   ".sc-1c0xbiw-9 > p:nth-child(1)",
	Description: ".iszfik-2>p:nth-child(1)",
	// MyBook's publisher value is unreliable - it can be public domain
	// Publisher:   ".hTmOzC > a:nth-child(1) > div:nth-child(1) > div:nth-child(1)",
	// FIX: this returns nil
	CoverImage:         ".kkiIwl > picture:nth-child(1) > img:nth-child(2)",
	CoverImageProperty: "srcset",
	Quirks: map[string]func(book *types.TempBook){
		"fixImageLink": func(book *types.TempBook) {
			// get the higher quality version
			split := strings.Split(book.ImageLink, ", ")[0]
			book.ImageLink = strings.Split(split, " ")[0]
		},
		"extractYear": func(book *types.TempBook) {
			split := strings.Split(book.Year, " ")
			book.Year = split[len(split)-1]
		},
		"extractPageCount": func(book *types.TempBook) {
			book.PageCount = strings.Split(book.PageCount, " ")[0]
		},
	},
}
