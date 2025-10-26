package main

import (
	"context"
	"housebook/api"
	"housebook/db"
	"log"
	"sync"
)

const (
	API_HOSTNAME = "0.0.0.0"
	DB_HOSTNAME  = "db"
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
		if err := api.Start(API_HOSTNAME, 3000); err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	go func() {
		if err := db.Start(DB_HOSTNAME, 5432); err != nil {
			log.Println(err)
		}
		wg.Done()
	}()
	wg.Wait()
}
