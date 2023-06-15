package users

import (
	"bench/middleware"

	"github.com/go-chi/chi/v5"
)

func New(r chi.Router) {
	r.With(
		middleware.WithAuth(),
	).Get(
		"/users/me",
		GetMe(),
	)

	r.Post(
		"/users",
		Create(),
	)

	r.With(
		middleware.WithAuth(),
	).Delete(
		"/users/me",
		Delete(),
	)
}
