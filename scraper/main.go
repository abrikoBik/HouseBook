package main

import (
	"context"
	"errors"
	"fmt"
	"housebook_scraper/consts"
	"housebook_scraper/handlers"
	"net"
	"net/http"
	"os"

	"github.com/rs/zerolog"
)

// func Scraper(url string, handler func(c *colly.Collector, url string) Book) Book {
// 	c := colly.NewCollector(colly.AllowedDomains("mybook.ru", "www.litres.ru", "www.labirint.ru"))

// 	c.OnRequest(func(r *colly.Request) {
// 		log.Println("Visiting: ", r.URL)
// 	})

// 	c.OnError(func(_ *colly.Response, err error) {
// 		log.Println("Something went wrong: ", err)
// 	})

// 	c.OnResponse(func(r *colly.Response) {
// 		log.Println("Page visited: ", r.Request.URL)
// 	})

// 	c.OnScraped(func(r *colly.Response) {
// 		log.Println(r.Request.URL, " scraped!")
// 	})

// 	res := handler(c, url)

// 	return res
// }

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	ctx := logger.WithContext(context.Background())
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", consts.Hostname, consts.Port))
	if err != nil {
		logger.Panic().Err(err).Msg("panicking at net.Listen")
	}
	defer listener.Close()
	server := http.Server{
		Addr:    listener.Addr().String(),
		Handler: handlers.Config(ctx),
	}
	logger.Info().Str("hostname", consts.Hostname).Int("port", consts.Port).Msg("serving")
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		logger.Panic().Err(err).Msg("panicking at server.Serve")
	}
	defer server.Shutdown(ctx)
}
