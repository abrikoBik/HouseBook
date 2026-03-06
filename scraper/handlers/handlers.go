// Package handlers contains all HTTP handler functions.
package handlers

import (
	"context"
	"housebook_scraper/consts"
	"housebook_scraper/sites"
	"net/http"

	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog"
)

var (
	ctx       context.Context
	logger    *zerolog.Logger
	collector *colly.Collector
)

type Request struct {
	URL string
}

func Config(c context.Context) *http.ServeMux {
	ctx = c
	logger = zerolog.Ctx(ctx)
	collector = sites.Config(ctx, consts.AllowedDomains)
	router := http.NewServeMux()
	router.HandleFunc("GET /health", Health)
	router.HandleFunc("POST /mybook", BuildScraper(collector, sites.ScrapeMyBook))
	router.HandleFunc("POST /labirint", BuildScraper(collector, sites.ScrapeLabirint))
	router.HandleFunc("POST /litres", BuildScraper(collector, sites.ScrapeLitres))
	return router
}

func Health(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
}
