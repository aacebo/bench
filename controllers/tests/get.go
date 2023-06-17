package tests

import (
	"net/http"

	"bench/models/problems"
	model "bench/models/tests"
	"bench/response"
)

func Get() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		problem := r.Context().Value("problem").(*problems.Problem)
		tests := model.GetByProblemID(*problem.ID)
		response.New(tests).Do(w, r)
	}
}
