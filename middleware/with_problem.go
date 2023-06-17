package middleware

import (
	"context"
	"net/http"

	model "bench/models/problems"
	"bench/response"

	"github.com/go-chi/chi/v5"
)

func WithProblem() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			problem := model.GetByName(chi.URLParam(r, "problem"))

			if problem == nil {
				response.NotFound("problem").Do(w, r)
				return
			}

			ctx = context.WithValue(ctx, "problem", problem)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
