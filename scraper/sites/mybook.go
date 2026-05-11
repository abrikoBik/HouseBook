package sites

import "housebook_scraper/consts"

var MyBook = Site{
	Domain:             consts.MyBookDomain,
	Root:               ".iDfqxd",
	Title:              "h1",
	Author:             ".sc-1c0xbiw-8>div:nth-child(1)>a:nth-child(1)>div:nth-child(1)>div:nth-child(1)",
	ISBN:               "div.iszfik-15:nth-child(2)>dl:nth-child(1)>dd:nth-child(2)",
	Genre:              "div.sc-1sg8rha-0:nth-child(2)>div:nth-child(1)>div:nth-child(1)>a:nth-child(1)",
	Year:               "html body div#__next div.sc-121k7ak-0.gpMNuT section.hf8eib-0.acq6ib-2.hYqYfq.gixdCj div.ant-row.ant-row-center.acq6ib-0.iONKLl div.ant-col.acq6ib-4.iDfqxd.ant-col-xs-12.ant-col-xs-order-3.ant-col-xl-6 div.m4n24q-0.feXlWY div.iszfik-0.fYTRQF div.ant-row.ant-row-center.iszfik-1.cFmLKi div.ant-col.ant-col-xs-11.ant-col-md-8.ant-col-xl-12 div.m4n24q-0.jAyNjF div.iszfik-2.NZYdY div.iszfik-3.gqSJbp div.m4n24q-0.hJyrxa div.iszfik-12.clyKYY div.iszfik-14.iSnZQd div.iszfik-15.BerVK dl.iszfik-16.dhqFsU dd.iszfik-18.iEusfO",
	PageCount:          ".sc-1c0xbiw-9>p:nth-child(1)",
	Description:        ".iszfik-2>p:nth-child(1)",
	CoverImage:         ".kkiIwl>picture:nth-child(1)>source:nth-child(1)",
	CoverImageProperty: "srcset",
}
