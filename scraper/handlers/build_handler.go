package handlers

import (
	"bytes"
	"encoding/json"
	"errors"
	"housebook_scraper/sites"
	"housebook_scraper/types"
	"io"
	"mime/multipart"
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
		// TODO: clean up
		if book.ImageLink != "" {
			logger.Info().Str("url", book.ImageLink).Msg("downloading")
			result, err := http.Get(book.ImageLink)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response.Error().Err(err).Send()
				return
			}
			defer result.Body.Close()
			logger.Info().Str("url", book.ImageLink).Msg("downloaded")

			logger.Info().Msg("uploading cover")
			buf := &bytes.Buffer{}
			mpw := multipart.NewWriter(buf)
			imageWriter, err := mpw.CreateFormFile("image", "cover")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response.Error().Err(err).Send()
				return
			}
			if _, err = io.Copy(imageWriter, result.Body); err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response.Error().Err(err).Send()
				return
			}
			mpw.Close()
			client := &http.Client{}
			req, err := http.NewRequest("POST", "http://storage:3001/", buf)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response.Error().Err(err).Send()
				return
			}
			req.Header.Add("Content-Type", mpw.FormDataContentType())
			res, err := client.Do(req)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				response.Error().Err(err).Send()
				return
			}
			defer res.Body.Close()
			logger.Info().Msg("uploaded cover")

			if res.StatusCode == http.StatusCreated {
				var body types.CDNResponse
				if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					response.Error().Err(err).Send()
					return
				}
				logger.Info().Str("imageUUID", body.ID)
				book.ImageUUID = body.ID
			} else {
				w.WriteHeader(http.StatusInternalServerError)
				var body types.CDNError
				if err := json.NewDecoder(res.Body).Decode(&body); err != nil {
					response.Error().Err(err).Send()
					return
				}
				err := errors.New(body.Error)
				logger.Error().Err(err).Send()
				response.Error().Err(err).Send()
				return
			}
		}
		if err := json.NewEncoder(w).Encode(book); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			response.Error().Err(err).Send()
			return
		}
	}
}
