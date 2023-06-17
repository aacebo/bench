package tests

import (
	"net/http"

	"bench/models/problems"
	model "bench/models/tests"
	"bench/response"
	schema "bench/schemas/tests"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		problem := r.Context().Value("problem").(*problems.Problem)
		body, errs := schema.NewCreate(r)

		if errs.Len() > 0 {
			response.BadReqest(errs).Do(w, r)
			return
		}

		test := model.New(
			*problem.ID,
			*body.Input,
			*body.Output,
		)

		test.Save()
		response.New(test).Status(201).Do(w, r)
	}
}
