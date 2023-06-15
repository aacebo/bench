package users

import "github.com/go-chi/chi/v5"

func New(r chi.Router) {
	r.Post(
		"/users",
		Create(),
	)
}
