package languages

import (
	"net/http"

	"bench/container"
	model "bench/models/languages"
	"bench/response"
)

func Run() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_ = r.Context().Value("lang").(*model.Language)
		c, err := container.New()

		if err != nil {
			response.Internal(err.Error()).Do(w, r)
			return
		}

		response.New(c).Do(w, r)
	}
}
