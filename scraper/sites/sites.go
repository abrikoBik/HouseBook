// Package sites contains scrapers for supported websites.
package sites

import (
	"context"

	"github.com/gocolly/colly/v2"
	"github.com/rs/zerolog"
)

func Config(ctx context.Context, domains []string) *colly.Collector {
	logger := zerolog.Ctx(ctx)
	collector := colly.NewCollector(
		colly.StdlibContext(ctx),
		colly.AllowedDomains(domains...),
		colly.AllowURLRevisit(),
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
	return collector
}
