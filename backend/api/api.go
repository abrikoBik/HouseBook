// Package api contains all API microservice properties, as well as subpackages for HTTP handlers and endpoints.
package api

import (
	"context"
	"errors"
	"fmt"
	"housebook/api/endpoints"
	"net"
	"net/http"

	"github.com/go-chi/chi/v5"
)

var (
	Context  context.Context
	Listener net.Listener
	Server   http.Server
	Router   chi.Router
	err      error
)

func Start(ctx context.Context, hostname string, port uint16, callback func()) {
	Context = ctx
	Listener, err = net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		panic(err)
	}
	defer Listener.Close()
	Server.Addr = Listener.Addr().String()

	Router = chi.NewRouter()
	Server.Handler = Router
	endpoints.Register(Router)

	callback()
	if err := Server.Serve(Listener); !errors.Is(err, http.ErrServerClosed) {
		panic(err)
	}
	defer Server.Shutdown(Context)
}

func Stop(callback func()) {
	Server.Close()
	callback()
}
