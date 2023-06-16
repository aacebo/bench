package problems

import (
	"bench/middleware"

	"github.com/go-chi/chi/v5"
)

func New(r chi.Router) {
	r.Get(
		"/problems",
		Get(),
	)

	r.With(
		middleware.WithAuth(),
	).Post(
		"/problems",
		Create(),
	)
}
