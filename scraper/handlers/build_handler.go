package handlers

import (
	"encoding/json"
	"housebook_scraper/sites"
	"net/http"
	"net/url"

	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog"
)

func BuildHandler(scraper sites.Site) func(http.ResponseWriter, *http.Request) {
	collector := colly.NewCollector(
		colly.StdlibContext(ctx),
		colly.AllowedDomains(scraper.Domain),
		colly.AllowURLRevisit(),
		colly.Async(true),
	)
	collector.OnError(func(r *colly.Response, err error) {
		logger.Info().Err(err).Send()
	})
	collector.OnRequest(func(r *colly.Request) {
		logger.Info().Str("url", r.URL.String()).Msg("scraping")
	})
	collector.OnScraped(func(r *colly.Response) {
		logger.Info().Str("url", r.Request.URL.String()).Msg("scraped")
	})
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
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
		if err := json.NewEncoder(w).Encode(book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response.Error().Err(err).Send()
			return
		}
	}
}
