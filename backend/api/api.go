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
)

func Start(hostname string, port uint16) error {
	if Context == nil {
		return errors.New("context is nil")
	}
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		return err
	}
	defer Listener.Close()
	Server.Addr = Listener.Addr().String()

	Router = chi.NewRouter()
	Server.Handler = Router
	endpoints.Register(Router)

	if err := Server.Serve(Listener); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	defer Server.Shutdown(Context)
	return nil
}
