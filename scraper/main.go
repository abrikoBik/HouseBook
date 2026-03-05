package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/gocolly/colly/v2"
)

type Book struct {
	title     string
	author    string
	ISBN      string
	genre     string
	language  string
	publisher string

	coverType   string
	pageCount   int
	description string
	color       bool
}

func ScrapeMyBook(c *colly.Collector, url string) Book {
	res := Book{}

	c.OnHTML(".iDfqxd", func(e *colly.HTMLElement) {
		res.title = e.ChildText("h1")
		res.author = e.ChildText(".eFHndH>.bkolKJ>a>.fZFkVd>.jVKkXg")
		res.ISBN = e.ChildText(".iSnZQd div:nth-child(2)>dl:nth-child(1)>dd")
		res.genre = e.ChildText(".cFmLKi>div:nth-child(2)>.gHinNz>div:nth-child(1)>div>a")
		// language - not found on MyBook.ru
		// publisher - not found on MyBook.ru

		// cover_type - not found on MyBook.ru
		pageCountString := e.ChildText(".eSjGMZ p:nth-child(1)")
		res.pageCount, _ = strconv.Atoi(pageCountString[0:strings.Index(pageCountString, " ")])
		res.description = e.ChildText(".gAFRve>p")
		// color - not found on MyBook.ru

	})

	c.Visit(url)

	return res
}

func ScrapeLitres(c *colly.Collector, url string) Book {
	res := Book{}

	c.OnHTML("body", func(e *colly.HTMLElement) {
		res.title = e.ChildText("h1")
		res.author = e.ChildText("._86af713b>div>span")
		res.ISBN = e.ChildText("._9d4e17ed>div:nth-child(5)>span")
		res.genre = e.ChildText(".aa63d864>div:nth-child(1)>.f12c0b54")
		// language - not found on Litres.ru
		res.publisher = e.ChildText("._9d4e17ed>div:nth-child(6)>a")

		// cover_type - not found on Litres.ru
		pageCountString := e.ChildText("._9d4e17ed>div:nth-child(4)>span")
		res.pageCount, _ = strconv.Atoi(pageCountString[0:strings.Index(pageCountString, " ")])
		res.description = e.ChildText(".ac83cc29>div:nth-child(1)")
		// color - not found on Litres.ru
	})

	c.Visit(url)

	return res
}

func ScrapeLabirint(c *colly.Collector, url string) Book {
	res := Book{}

	c.OnHTML("body", func(e *colly.HTMLElement) {
		res.title = e.ChildText("h1")
		res.author = e.ChildText("._features_ctofl_78>div:nth-child(1)>div:nth-child(2)>a")
		res.ISBN = e.ChildText("._right_u86in_12>div:nth-child(2)>.text-black>span")
		// genre - not found on Labirint.ru
		res.language = e.ChildText("._right_u86in_12>div:nth-child(3)>.text-black>span")
		res.publisher = e.ChildText("._left_u86in_12>div:nth-child(4)>div:nth-child(2)>a")

		res.coverType = e.ChildText("._right_u86in_12>div:nth-child(7)>.text-black>span")
		pageCountString := e.ChildText("._right_u86in_12>div:nth-child(4)>.text-black>span:nth-child(1)")
		pageCountSlice := strings.Split(pageCountString, " ")
		if len(pageCountSlice) == 0 {
			log.Println("No page count found!")
			return
		}
		var err error
		if res.pageCount, err = strconv.Atoi(pageCountSlice[0]); err != nil {
			log.Println("Ero", err)
			return
		}
		res.description = e.ChildText("._wrapper_1rsml_1>div:nth-child(1)>div")
		// color - not found on Labirint.ru
	})

	c.Visit(url)

	return res
}

func Scraper(url string, handler func(c *colly.Collector, url string) Book) Book {
	c := colly.NewCollector(colly.AllowedDomains("mybook.ru", "www.litres.ru", "www.labirint.ru"))

	c.OnRequest(func(r *colly.Request) {
		log.Println("Visiting: ", r.URL)
	})

	c.OnError(func(_ *colly.Response, err error) {
		log.Println("Something went wrong: ", err)
	})

	c.OnResponse(func(r *colly.Response) {
		log.Println("Page visited: ", r.Request.URL)
	})

	c.OnScraped(func(r *colly.Response) {
		log.Println(r.Request.URL, " scraped!")
	})

	res := handler(c, url)

	return res
}

func main() {
	book := Scraper("https://mybook.ru/author/fedor-mihajlovich-dostoevskij/prestuplenie-i-nakazanie/", ScrapeMyBook)

	var b strings.Builder

	b.WriteString("Название книги: ")
	b.WriteString(book.title)
	b.WriteRune('\n')
	b.WriteString("Автор: ")
	b.WriteString(book.author)
	b.WriteRune('\n')
	b.WriteString("Жанр: ")
	b.WriteString(book.genre)
	b.WriteRune('\n')
	b.WriteString("Издатель: ")
	b.WriteString(book.publisher)
	b.WriteRune('\n')
	b.WriteString("ISBN: ")
	b.WriteString(book.ISBN)
	b.WriteRune('\n')
	b.WriteString("Число страниц: ")
	b.WriteString(strconv.Itoa(book.pageCount))
	b.WriteRune('\n')
	b.WriteString("Язык: ")
	b.WriteString(book.language)
	b.WriteRune('\n')
	b.WriteString("Тип обложки: ")
	b.WriteString(book.coverType)
	b.WriteRune('\n')
	b.WriteString("Описание книги:")
	b.WriteRune('\n')
	b.WriteString(book.description)

	fmt.Print(b.String())

	//fmt.Println("Название книги:", book.title)
	//fmt.Println("Автор книги:", book.author)
	//fmt.Println("Жанр:", book.genre)
	//fmt.Println("Издатель:", book.publisher)
	//fmt.Println("ISBN:", book.ISBN)
	//fmt.Println("Число страниц:", book.page_count)
	//fmt.Println("Язык:", book.language)
	//fmt.Println("Тип обложки:", book.cover_type)
	//fmt.Println("Описание книги:\n", book.description)
}
