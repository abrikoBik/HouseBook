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
		wg        sync.WaitGroup
		exit      = make(chan os.Signal, 1)
		errLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stderr})
		outLogger = zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout})
		ctx       = context.WithValue(
			context.WithValue(context.Background(), consts.OutLogger, outLogger),
			consts.ErrLogger, errLogger)
	)
	signal.Notify(exit, syscall.SIGTERM)
	wg.Add(consts.ServiceCount)
	go api.Start(ctx, consts.APIHostname, consts.APIPort, wg.Done)
	go database.Start(ctx, consts.DBHostname, consts.DBPort, wg.Done)
	wg.Wait()
	outLogger.Info().Int("serviceCount", consts.ServiceCount).Msg("starting")

	<-exit
	outLogger.Info().Int("serviceCount", consts.ServiceCount).Msg("stopping")
	wg.Add(consts.ServiceCount)
	go api.Stop(wg.Done)
	go database.Stop(wg.Done)
	wg.Wait()
}
