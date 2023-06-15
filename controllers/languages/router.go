package languages

import (
	"bench/middleware"

	"github.com/go-chi/chi/v5"
)

func New(r chi.Router) {
	r.Get(
		"/languages",
		Get(),
	)

	r.With(
		middleware.WithAuth(),
		middleware.WithLanguage(),
	).Post(
		"/languages/{lang}",
		Run(),
	)
}
