package middleware

import (
	"context"
	"net/http"

	model "bench/models/languages"
	"bench/response"

	"github.com/go-chi/chi/v5"
)

func WithLanguage() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			lang := model.GetByName(chi.URLParam(r, "lang"))

			if lang == nil {
				response.NotFound("language").Do(w, r)
				return
			}

			ctx = context.WithValue(ctx, "lang", lang)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
