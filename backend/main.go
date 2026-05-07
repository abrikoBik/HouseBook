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

	wg.Add(consts.ServiceCount)
	dbConf := database.Conf{
		Context: logger.WithContext(ctx),
		// TODO: implement env module
		Username: os.Getenv("POSTGRES_USER"),
		Password: os.Getenv("POSTGRES_PASSWORD"),
		Database: os.Getenv("POSTGRES_DB"),
		Hostname: "db",
		Port:     5432,
	}
	apiConf := api.Conf{
		Context:  logger.WithContext(ctx),
		JWTKey:   []byte(os.Getenv("JWT_KEY")),
		Hostname: "0.0.0.0",
		Port:     3000,
	}
	go database.Config(dbConf, wg.Done)
	go api.Config(apiConf, wg.Done)
	wg.Wait()
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("configured")

	wg.Add(consts.ServiceCount)
	go database.Start(consts.DBHostname, consts.DBPort, wg.Done)
	go api.Start(wg.Done)
	wg.Wait()
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("started")

	signal.Notify(exit, syscall.SIGTERM)
	<-exit
	wg.Add(consts.ServiceCount)
	go database.Stop(wg.Done)
	go api.Stop(wg.Done)
	wg.Wait()
	logger.Info().Int("serviceCount", consts.ServiceCount).Msg("stopped")
}
