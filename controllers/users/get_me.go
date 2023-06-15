package users

import (
	"net/http"

	model "bench/models/users"
	"bench/response"
)

func GetMe() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		me := r.Context().Value("me").(*model.User)
		response.New(me).Do(w, r)
	}
}
