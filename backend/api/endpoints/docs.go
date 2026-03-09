package endpoints

import (
	"fmt"
	"housebook/consts"
	_ "housebook/docs"
	"net/http"

	"github.com/go-chi/chi/v5"
	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func routeDocs(router chi.Router) {
	router.Route(consts.DocsEndpoint, func(r chi.Router) {
		r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
			http.ServeFile(w, r, "/app/swagger.json")
		})
		r.Get("/*", httpSwagger.Handler(
			httpSwagger.URL(fmt.Sprintf("/api%s/swagger.json", consts.DocsEndpoint)),
		))
	})
}
