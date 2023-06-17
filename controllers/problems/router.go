package problems

import (
	"bench/middleware"
	"bench/models/users"

	"github.com/go-chi/chi/v5"
)

func New(r chi.Router) {
	r.Get(
		"/problems",
		Get(),
	)

	r.With(
		middleware.WithAuth(users.ADMIN),
	).Post(
		"/problems",
		Create(),
	)

	r.With(
		middleware.WithAuth(users.ADMIN),
		middleware.WithProblem(),
	).Patch(
		"/problems/{problem}",
		Update(),
	)
}
