package main

import (
	"context"
	"housebook/api"
	"housebook/consts"
	"housebook/database"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/rs/zerolog"
)

func main() {
	var (
		wg     sync.WaitGroup
		exit   = make(chan os.Signal, 1)
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).
			With().Timestamp().
			Logger()
		ctx = context.Background()
	)
	ctx = logger.WithContext(ctx)
	signal.Notify(exit, syscall.SIGTERM)
	wg.Add(consts.ServiceCount)
	go database.Config(ctx, wg.Done)
	go api.Config(ctx, wg.Done)
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("configured")

	wg.Add(consts.ServiceCount)
	go database.Start(consts.DBHostname, consts.DBPort, wg.Done)
	go api.Start(consts.APIHostname, consts.APIPort, wg.Done)
	wg.Wait()
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("started")

	<-exit
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("stopping")
	wg.Add(consts.ServiceCount)
	go database.Stop(wg.Done)
	go api.Stop(wg.Done)
	wg.Wait()
}
