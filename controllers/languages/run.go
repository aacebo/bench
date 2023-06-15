package languages

import (
	"net/http"

	model "bench/models/languages"
	"bench/response"
)

func Run() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		lang := r.Context().Value("lang").(*model.Language)
		response.New(lang).Do(w, r)
	}
}
