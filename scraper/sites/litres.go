package sites

import "housebook_scraper/consts"

var Litres = Site{
	Domain:             consts.LitresDomain,
	Root:               "body",
	Title:              "h1",
	Author:             "li.Breadcrumbs_breadcrumbs__item__zPK82:nth-child(3) > a:nth-child(1) > span:nth-child(2)",
	ISBN:               ".CharacteristicsBlock_characteristics__1XOas>div:nth-child(6)>span",
	Genre:              "li.Breadcrumbs_breadcrumbs__item__zPK82:nth-child(2) > a:nth-child(1) > span:nth-child(1)",
	Publisher:          ".CharacteristicsBlock_characteristics__1XOas>div:nth-child(8)>a",
	Year:               "div.CharacteristicsBlock_characteristic__4pi7v:nth-child(4) > span:nth-child(2)",
	PageCount:          ".CharacteristicsBlock_characteristics__1XOas>div:nth-child(5)>span",
	Description:        ".Truncate_expandedModern__c4VSi>div:nth-child(1)>p:nth-child(2)",
	CoverImage:         ".BookCover_img__ncTI5",
	CoverImageProperty: "src",
}
