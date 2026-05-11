package sites

import "housebook_scraper/consts"

var Litres = Site{
	Domain: consts.LitresDomain,
	Root:   "body",
	Title:  "h1",
	// FIX: Litres refuses to return author information
	Author:             ".BookAuthor_author__name__ORVlE>div:nth-child(1)>div:nth-child(1)>span:nth-child(1)",
	ISBN:               ".CharacteristicsBlock_characteristics__1XOas>div:nth-child(6)>span",
	Genre:              ".BookGenresAndTags_genresList__rd8vU>div:nth-child(1)>a",
	Publisher:          ".CharacteristicsBlock_characteristics__1XOas>div:nth-child(8)>a",
	Year:               ".CharacteristicsBlock_characteristics__1XOas>div:nth-child(3)>span",
	PageCount:          ".CharacteristicsBlock_characteristics__1XOas>div:nth-child(5)>span",
	Description:        ".Truncate_expandedModern__c4VSi>div:nth-child(1)>p:nth-child(2)",
	CoverImage:         ".BookCover_img__ncTI5",
	CoverImageProperty: "src",
}
