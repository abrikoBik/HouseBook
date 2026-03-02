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
	"github.com/go-chi/chi/v5/middleware"
	"github.com/rs/zerolog"
)

var (
	Context context.Context
	Logger  *zerolog.Logger

	addr     string
	listener net.Listener
	server   http.Server
	router   chi.Router
	err      error
)

type Conf struct {
	Context  context.Context
	Hostname string
	Port     uint16
}

func Config(config Conf, callback func()) {
	Context = config.Context
	Logger = zerolog.Ctx(Context)
	handlers.Config(Context)
	addr = fmt.Sprintf("%s:%d", config.Hostname, config.Port)
	router = chi.NewRouter()
	// TODO: add in production
	// router.Use(middleware.Compress(5))
	router.Use(middleware.Recoverer)
	endpoints.Register(router)
	callback()
}

// @title HouseBook API
// @version 1.0

// @license.name GNU General Public License
// @license.url https://www.gnu.org/licenses/gpl-3.0.en.html

// @host 0.0.0.0
// @BasePath /api

func Start(callback func()) {
	listener, err = net.Listen("tcp", addr)
	if err != nil {
		Logger.Panic().Err(err).Msg("panicking in API")
	}
	defer listener.Close()
	server.Addr = listener.Addr().String()
	server.Handler = router
	Logger.Info().Str("addr", addr).Msg("starting API server")
	callback()
	if err := server.Serve(listener); !errors.Is(err, http.ErrServerClosed) {
		Logger.Panic().Err(err).Msg("panicking in API")
	}
	defer server.Shutdown(Context)
}

func Stop(callback func()) {
	Logger.Info().Msg("stopping API server")
	server.Close()
	callback()
}
