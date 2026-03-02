// Package api contains all API microservice properties, as well as subpackages for HTTP handlers and endpoints.
package api

import (
	"context"
	"errors"
	"fmt"
	"housebook/api/endpoints"
	"housebook/api/handlers"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

var (
	Context  context.Context
	Listener net.Listener
	Server   http.Server
	Router   chi.Router
	Logger   *zerolog.Logger
	err      error
)

func Config(ctx context.Context, callback func()) {
	Context = ctx
	Logger = zerolog.Ctx(Context)
	handlers.Config(Context)
	callback()
}

func Start(hostname string, port uint16, callback func()) {
	Listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		Logger.Panic().Err(err).Msg("panicking in API")
	}
	defer Listener.Close()
	Server.Addr = Listener.Addr().String()

	Router = chi.NewRouter()
	Server.Handler = Router
	endpoints.Register(Router)

	Logger.Info().Str("hostname", hostname).Uint16("port", port).Msg("starting API server")
	callback()
	if err := Server.Serve(Listener); !errors.Is(err, http.ErrServerClosed) {
		Logger.Panic().Err(err).Msg("panicking in API")
	}
	defer Server.Shutdown(Context)
}

func Stop(callback func()) {
	Logger.Info().Msg("stopping API server")
	Server.Close()
	callback()
}
