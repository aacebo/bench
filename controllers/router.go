package controllers

import (
	"bench/controllers/languages"
	"bench/controllers/problems"
	"bench/controllers/sessions"
	"bench/controllers/tests"
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
		sessions.New(r)
		languages.New(r)
		problems.New(r)
		tests.New(r)
	})

	return r
}
