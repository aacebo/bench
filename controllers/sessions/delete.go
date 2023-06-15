package sessions

import (
	"net/http"

	model "bench/models/sessions"
	"bench/response"
)

func Delete() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value("session").(*model.Session)
		session.Delete()
		response.New(session).Do(w, r)
	}
}
