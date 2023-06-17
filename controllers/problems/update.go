package problems

import (
	"net/http"

	model "bench/models/problems"
	"bench/models/users"
	"bench/response"
	schema "bench/schemas/problems"
)

func Update() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		me := r.Context().Value("me").(*users.User)
		problem := r.Context().Value("problem").(*model.Problem)

		if *me.ID != *problem.CreatedById {
			response.Unauthorized().Do(w, r)
			return
		}

		body, errs := schema.NewUpdate(r)

		if errs.Len() > 0 {
			response.BadReqest(errs).Do(w, r)
			return
		}

		if body.Name != nil {
			existing := model.GetByName(*body.Name)

			if existing != nil && *existing.ID != *problem.ID {
				response.Conflict("name").Do(w, r)
				return
			}

			problem.Name = body.Name
		}

		if body.DisplayName != nil {
			problem.DisplayName = body.DisplayName
		}

		if body.Description != nil {
			problem.Description = body.Description
		}

		problem.Save()
		response.New(problem).Do(w, r)
	}
}
