package main

import (
	"context"
	"housebook/api"
	"housebook/db"
	"log"
	"sync"
)

const (
	apiHostname      = "0.0.0.0"
	databaseHostname = "db"
)

func init() {
	ctx := context.Background()
	api.Context = ctx
	db.Context = ctx
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		if err := api.Start(apiHostname, 3000); err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	go func() {
		if err := db.Start(databaseHostname, 5432); err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
