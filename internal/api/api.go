package api

import (
	v1 "dh-auth/internal/api/v1"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func NewHandler() http.Handler {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	r.Get("/healthz", func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusOK)
	})
	v1Handler := v1.NewHandler()
	r.Mount("/v1", v1.Routes(v1Handler))
	return r
}
