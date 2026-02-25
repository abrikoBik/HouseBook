package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"storage/consts"
	"storage/handlers"
)

func main() {
	ctx := context.Background()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", consts.Hostname, consts.Port))
	if err != nil {
		log.Panicln(err.Error())
	}
	defer listener.Close()
	server := http.Server{
		Addr:    listener.Addr().String(),
		Handler: handlers.Config(ctx),
	}
	log.Printf("listening at %s:%d\n", consts.Hostname, consts.Port)
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		log.Panicln(err.Error())
	}
	defer server.Shutdown(ctx)
}
