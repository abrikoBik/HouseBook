package main

import (
	"context"
	"errors"
	"fmt"
	"net"
	"net/http"
	"os"
	"storage/consts"
	"storage/handlers"

	"github.com/rs/zerolog"
)

func main() {
	logger := zerolog.New(zerolog.ConsoleWriter{Out: os.Stdout}).With().Timestamp().Logger()
	ctx := logger.WithContext(context.Background())
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", consts.Hostname, consts.Port))
	if err != nil {
		logger.Panic().Err(err).Msg("panicking at net.Listen")
	}
	defer listener.Close()
	server := http.Server{
		Addr:    listener.Addr().String(),
		Handler: handlers.Config(ctx),
	}
	logger.Info().Str("hostname", consts.Hostname).Int("port", consts.Port).Msg("serving")
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		logger.Panic().Err(err).Msg("panicking at server.Serve")
	}
	defer server.Shutdown(ctx)
}
