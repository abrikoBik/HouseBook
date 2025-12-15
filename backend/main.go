package main

import (
	"context"
	"housebook/api"
	"housebook/database"
)

const (
	apiHostname      = "0.0.0.0"
	databaseHostname = "db"
)

func init() {
	ctx := context.Background()
	database.Context = ctx
	api.Context = ctx
}

func main() {
	if err := database.Start(databaseHostname, 5432); err != nil {
		panic(err)
	}
	if err := api.Start(apiHostname, 3000); err != nil {
		panic(err)
	}
}
