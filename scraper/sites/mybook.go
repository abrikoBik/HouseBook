package sites

import "housebook_scraper/consts"

var MyBook = Site{
	Domain:      consts.MyBookDomain,
	Root:        ".iDfqxd",
	Title:       "h1",
	Author:      ".eFHndH>.bkolKJ>a>.fZFkVd>.jVKkXg",
	ISBN:        ".iSnZQd div:nth-child(2)>dl:nth-child(1)>dd",
	Genre:       ".cFmLKi>div:nth-child(2)>.gHinNz>div:nth-child(1)>div>a",
	Year:        ".iSnZQd div:nth-child(1)>dl:nth-child(3)>dd",
	PageCount:   ".eSjGMZ p:nth-child(1)",
	Description: ".gAFRve>p",
	CoverImage:  ".idoTYd>.kkiIwl>picture", // TODO: fix
}
