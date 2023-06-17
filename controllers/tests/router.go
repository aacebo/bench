package tests

import (
	"bench/middleware"

	"github.com/go-chi/chi/v5"
)

func New(r chi.Router) {
	r.With(
		middleware.WithProblem(),
	).Get(
		"/problems/{problem}/tests",
		Get(),
	)
}
