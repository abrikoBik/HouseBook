package sites

import (
	"housebook_scraper/consts"
	"housebook_scraper/types"
	"strings"
)

var Labirint = Site{
	Domain:             consts.LabirintDomain,
	Root:               "body",
	Title:              "h1",
	Author:             "._features_ctofl_78>div:nth-child(1)>div:nth-child(2)>a",
	ISBN:               "._right_u86in_12>div:nth-child(2)>.text-black>span",
	Language:           "._right_u86in_12>div:nth-child(3)>.text-black>span",
	Publisher:          "._left_u86in_12>div:nth-child(4)>div:nth-child(2)>a",
	Year:               "._left_u86in_12>div:nth-child(4)>div:nth-child(2)>span:last-of-type",
	CoverType:          "._right_u86in_12>div:nth-child(7)>.text-black>span",
	PageCount:          "._right_u86in_12>div:nth-child(4)>.text-black>span:nth-child(1)",
	Description:        "._wrapper_1rsml_1>div:nth-child(1)>div",
	Genre:              "div.relative:nth-child(9)>div:nth-child(1)>span:last-child>a:nth-child(1)>span:nth-child(1)",
	CoverImage:         "._image_1qke2_7",
	CoverImageProperty: "srcset",
	Quirks: map[string]func(book *types.TempBook){
		"fixTitle": func(book *types.TempBook) {
			split := strings.Split(book.Title, ": ")
			split = split[:len(split)-1]
			book.Title = strings.Join(split, ": ")
		},
		"fixImageLink": func(book *types.TempBook) {
			book.ImageLink = strings.Split(book.ImageLink, " ")[0]
		},
	},
}
