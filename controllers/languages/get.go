package languages

import (
	"net/http"

	model "bench/models/languages"
	"bench/response"
)

func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		langs := model.Get()
		response.New(langs).Do(w, r)
	}
}
