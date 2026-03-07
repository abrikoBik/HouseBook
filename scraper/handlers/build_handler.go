package handlers

import (
	"encoding/json"
	"housebook_scraper/sites"
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog"
)

func BuildHandler(collector *colly.Collector, scraper sites.Scraper) func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var (
			requestData Request
			requestURL  *url.URL
			err         error
			response    = zerolog.New(w)
		)
		if err := json.NewDecoder(r.Body).Decode(&requestData); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Error().Err(err).Send()
			return
		}
		if requestURL, err = url.ParseRequestURI(requestData.URL); err != nil {
			w.WriteHeader(http.StatusBadRequest)
			response.Error().Err(err).Send()
			return
		}
		book, err := scraper.Scrape(collector, requestURL)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response.Error().Err(err).Send()
			return
		}
		response.Log().Interface("book", book).Send()
	}
}
