package api

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

var router = chi.NewRouter()

// Router returns the api router
func Router() http.Handler {
	router.Use(middleware.Logger)
	router.Use(middleware.Recoverer)

	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			Status: http.StatusOK,
			Data:   "Welcome to <no value>",
		}
		resp.ServerJSON(w)
	})

	router.NotFound(func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			Status: http.StatusOK,
			Data:   "route not found",
		}
		resp.ServerJSON(w)
	})

	router.MethodNotAllowed(func(w http.ResponseWriter, r *http.Request) {
		resp := response{
			Status: http.StatusOK,
			Data:   "method not allowed",
		}
		resp.ServerJSON(w)
	})

	registerRoutes()

	return router
}

func registerRoutes() {
	router.Route("/v1", func(r chi.Router) {
		r.Get("/", index)
	})
}
