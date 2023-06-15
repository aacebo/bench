package controllers

import (
	"bench/controllers/users"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/render"
)

func New() *chi.Mux {
	now := time.Now()
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		render.JSON(w, r, map[string]interface{}{
			"uptime": now.UnixMilli(),
		})
	})

	r.Route("/v1", func(r chi.Router) {
		users.New(r)
	})

	return r
}
