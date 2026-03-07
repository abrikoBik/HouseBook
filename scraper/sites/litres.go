package sites

import "housebook_scraper/consts"

var Litres = Site{
	Domain:      consts.LitresDomain,
	Root:        "body",
	Title:       "h1",
	Author:      "._86af713b>div>span",
	ISBN:        "._9d4e17ed>div:nth-child(6)>span",
	Genre:       ".aa63d864>div:nth-child(1)>.f12c0b54",
	Publisher:   "._9d4e17ed>div:nth-child(8)>a",
	Year:        "._9d4e17ed>div:nth-child(3)>span",
	PageCount:   "._9d4e17ed>div:nth-child(5)>span",
	Description: ".ac83cc29>div:nth-child(1)",
}
