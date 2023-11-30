package midlewares

import (
	"context"
	"diplomPlugService/internal/models"
	"net/http"

	"github.com/golang-jwt/jwt"
)

// func MustAuthenticate(h http.HandlerFunc) http.HandlerFunc {
// 	return func(w http.ResponseWriter, r *http.Request) {
// 		user, err := getCurrentUser(r)
// 		if err != nil {
// 			// write error code then return
// 		}
// 		ctx := context.WithValue(r.Context(), someKey, someValue)
// 		h.ServeHTTP(w, r.WithContext(ctx))
// 	}
// }

func CheckAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		println(authHeader)
		var infoFromToken models.JwtClaims

		_, err := jwt.ParseWithClaims(authHeader, &infoFromToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret word"), nil
		})
		if err != nil {
			w.WriteHeader(401)
			return
		}
		if infoFromToken.Type != "access" {
			w.WriteHeader(401)
			return
		}
		ctx := context.WithValue(r.Context(), "user", infoFromToken.User)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
