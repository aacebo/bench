package tests

import (
	"bench/middleware"
	"bench/models/users"

	"github.com/go-chi/chi/v5"
)

func New(r chi.Router) {
	r.With(
		middleware.WithProblem(),
	).Get(
		"/problems/{problem}/tests",
		Get(),
	)

	r.With(
		middleware.WithAuth(users.ADMIN),
		middleware.WithProblem(),
	).Post(
		"/problems/{problem}/tests",
		Create(),
	)
}
