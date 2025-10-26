// Package api contains all API microservice properties, as well as subpackages for HTTP handlers and endpoints.
package api

import (
	"context"
	"errors"
	"fmt"
	"housebook/api/endpoints"
	"net"
	"net/http"
)

var (
	Context  context.Context
	Listener net.Listener
	Server   http.Server
	Mux      *http.ServeMux
)

func Start(hostname string, port uint16) error {
	Listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", hostname, port))
	if err != nil {
		return err
	}
	defer Listener.Close()
	Server.Addr = Listener.Addr().String()

	Mux = http.NewServeMux()
	endpoints.Register(Mux)
	Server.Handler = Mux

	if err := Server.Serve(Listener); !errors.Is(err, http.ErrServerClosed) {
		return err
	}
	defer Server.Shutdown(Context)
	return nil
}
