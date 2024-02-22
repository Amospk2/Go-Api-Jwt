package middleware

import (
	"context"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	User string `json:"user"`
	Exp  string `json:"exp"`
}

func AuthenticationMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")

		accessToken := r.Header.Get("Bearrer")
		if accessToken == "" {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		parsedAccessToken, _ := jwt.ParseWithClaims(
			accessToken, jwt.MapClaims{},
			func(token *jwt.Token) (interface{}, error) {
				return []byte("secret"), nil
			},
		)

		if parsedAccessToken.Valid {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		ctx := context.WithValue(r.Context(), "user", parsedAccessToken.Claims)

		r = r.WithContext(ctx)

		next(w, r)
	})
}
