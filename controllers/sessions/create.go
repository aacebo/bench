package sessions

import (
	"net/http"
	"time"

	model "bench/models/sessions"
	users "bench/models/users"
	"bench/response"
	schema "bench/schemas/sessions"
	"bench/utils"

	"github.com/golang-jwt/jwt/v4"
)

func Create() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		body, errs := schema.NewCreate(r)

		if errs.Len() > 0 {
			response.BadReqest(errs).Do(w, r)
			return
		}

		user := users.GetByEmail(*body.Email)

		if user == nil || *body.Password != *user.Password {
			response.Unauthorized().Do(w, r)
			return
		}

		session := model.GetByUserID(*user.ID)

		if session == nil {
			session = model.New(*user.ID)
			session.Save()
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
			"sub": *session.ID,
			"iat": time.Now().Unix(),
		})

		tokenstr, _ := token.SignedString([]byte(utils.GetEnv("SECRET", "secret123!")))
		res := response.New(session).Status(201)
		res = res.SetMeta("token", tokenstr)
		res.Do(w, r)
	}
}
