package middleware

import (
	"fb-service/config"
	"fb-service/helper"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				response := map[string]string{"message": "Unauthorized"}
				helper.ResponseJson(w, http.StatusUnauthorized, response)
				return
			}
		}

		tokenString := c.Value

		claims := &config.JWTClaim{}
		//parsing jwt token
		token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
			return config.JWT_KEY, nil
		})

		if err != nil || !token.Valid {
			response := map[string]string{"message": "Unauthorized!"}
			helper.ResponseJson(w, http.StatusUnauthorized, response)
			return
		}

		next.ServeHTTP(w, r)
	})
}
