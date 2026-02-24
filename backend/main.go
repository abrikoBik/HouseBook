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
		logger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
		ctx    = context.Background()
	)
	ctx = logger.WithContext(ctx)
	signal.Notify(exit, syscall.SIGTERM)
	wg.Add(consts.ServiceCount)
	database.Start(ctx, consts.DBHostname, consts.DBPort, wg.Done)
	go api.Start(ctx, consts.APIHostname, consts.APIPort, wg.Done)
	wg.Wait()
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("started")

	<-exit
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("stopping")
	wg.Add(consts.ServiceCount)
	api.Stop(wg.Done)
	database.Stop(wg.Done)
	wg.Wait()
}
