package sessions

import (
	"bench/middleware"

	"github.com/go-chi/chi/v5"
)

func New(r chi.Router) {
	r.Post("/sessions", Create())
	r.With(middleware.WithAuth()).Delete("/sessions", Delete())
}
