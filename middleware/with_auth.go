package middleware

import (
	"bench/models/sessions"
	"bench/models/users"
	"bench/response"
	"bench/utils"
	"context"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/exp/slices"
)

func WithAuth(types ...users.Type) func(http.Handler) http.Handler {
	if len(types) == 0 {
		types = []users.Type{users.USER, users.ADMIN}
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := r.Context()
			token := r.Header.Get("x-bench-token")

			if token == "" {
				response.Unauthorized().Do(w, r)
				return
			}

			decoded, err := jwt.Parse(token, func(decoded *jwt.Token) (interface{}, error) {
				if _, ok := decoded.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("unexpected signing method: %v", decoded.Header["alg"])
				}

				return []byte(utils.GetEnv("SECRET", "secret123!")), nil
			})

			if err != nil {
				response.Unauthorized().Do(w, r)
				return
			}

			session := sessions.GetByID(decoded.Claims.(jwt.MapClaims)["sub"].(string))

			if session == nil {
				response.Unauthorized().Do(w, r)
				return
			}

			me := users.GetByID(*session.UserID)

			if me == nil || !slices.Contains(types, me.Type) {
				response.Unauthorized().Do(w, r)
				return
			}

			ctx = context.WithValue(ctx, "session", session)
			ctx = context.WithValue(ctx, "me", me)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
