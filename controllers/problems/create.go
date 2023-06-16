package problems

import (
	"net/http"

	model "bench/models/problems"
	"bench/models/users"
	"bench/response"
	schema "bench/schemas/problems"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		me := r.Context().Value("me").(*users.User)
		body, errs := schema.NewCreate(r)

		if errs.Len() > 0 {
			response.BadReqest(errs).Do(w, r)
			return
		}

		existing := model.GetByName(*body.Name)

		if existing != nil {
			response.Conflict("name").Do(w, r)
			return
		}

		problem := model.New(
			*body.Name,
			*body.DisplayName,
			*body.Description,
			*me.ID,
		)

		problem.Save()
		response.New(problem).Status(201).Do(w, r)
	}
}
