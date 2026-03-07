package sites

import "housebook_scraper/consts"

var Labirint = Site{
	Domain:      consts.LabirintDomain,
	Root:        "body",
	Title:       "h1",
	Author:      "._features_ctofl_78>div:nth-child(1)>div:nth-child(2)>a",
	ISBN:        "._right_u86in_12>div:nth-child(2)>.text-black>span",
	Language:    "._right_u86in_12>div:nth-child(3)>.text-black>span",
	Publisher:   "._left_u86in_12>div:nth-child(4)>div:nth-child(2)>a",
	CoverType:   "._right_u86in_12>div:nth-child(7)>.text-black>span",
	PageCount:   "._right_u86in_12>div:nth-child(4)>.text-black>span:nth-child(1)",
	Description: "._wrapper_1rsml_1>div:nth-child(1)>div",
}
