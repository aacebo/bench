package users

import (
	"net/http"

	model "bench/models/users"
	"bench/response"
	schema "bench/schemas/users"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, errs := schema.NewCreate(r)

		if errs.Len() > 0 {
			response.BadReqest(errs).Do(w, r)
			return
		}

		existing := model.GetByEmail(*body.Email)

		if existing != nil {
			response.Conflict("email").Do(w, r)
			return
		}

		user := model.New(*body.Name, *body.Email, *body.Password)
		user.Save()
		response.New(user).Status(201).Do(w, r)
	}
}
