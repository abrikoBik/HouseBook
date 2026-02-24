package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
)

const (
	hostname = "0.0.0.0"
	port     = 3001
)

func main() {
	ctx := context.Background()
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		log.Panicln(err.Error())
	}
	defer listener.Close()

	server := http.Server{
		Addr: listener.Addr().String(),
	}

	log.Printf("listening at %s:%d\n", hostname, port)
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		log.Panicln(err.Error())
	}
	defer server.Shutdown(ctx)
}
