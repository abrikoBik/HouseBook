package main

import (
	"context"
	"housebook/api"
	"housebook/database"
	"log"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

const (
	apiHostname      = "0.0.0.0"
	databaseHostname = "db"
	serviceCount     = 2
)

func main() {
	// TODO: proper logging solution
	var (
		wg   sync.WaitGroup
		ctx  = context.Background()
		exit = make(chan os.Signal, 1)
	)
	signal.Notify(exit, syscall.SIGTERM)
	wg.Add(serviceCount)
	go api.Start(ctx, apiHostname, 3000, wg.Done)
	go database.Start(ctx, databaseHostname, 5432, wg.Done)
	wg.Wait()
	log.Println(serviceCount, "services running")

	<-exit
	log.Println("SIGTERM received; stopping", serviceCount, "services")
	wg.Add(serviceCount)
	go api.Stop(wg.Done)
	go database.Stop(wg.Done)
	wg.Wait()
}
