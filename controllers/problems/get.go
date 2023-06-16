package problems

import (
	"net/http"

	model "bench/models/problems"
	"bench/response"
)

func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		problems := model.Get()
		response.New(problems).Do(w, r)
	}
}
