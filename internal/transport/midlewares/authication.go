package midlewares

import (
	"context"
	"diplomPlugService/internal/models"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt"
)

var routsWithoutAuthorization [2]string = [2]string{"/login", "/refresh"}

func CheckAuthorization(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for _, rout := range routsWithoutAuthorization {
			if r.URL.String() == rout {
				next.ServeHTTP(w, r)
				return
			}
		}
		authHeader := strings.Split(r.Header.Get("Authorization"), " ")
		if len(authHeader) != 2 {
			w.WriteHeader(401)
			w.Write([]byte("1invalid token"))
			return
		}

		tokenPrefix := authHeader[0]
		tokenString := authHeader[1]

		if tokenPrefix != "Bearer" {
			w.WriteHeader(401)
			w.Write([]byte("2invalid token"))
			return
		}

		var infoFromToken models.JwtClaims

		_, err := jwt.ParseWithClaims(tokenString, &infoFromToken, func(token *jwt.Token) (interface{}, error) {
			return []byte("secret word"), nil
		})
		if err != nil {
			w.WriteHeader(401)
			w.Write([]byte("3invalid token"))
			return
		}
		if infoFromToken.Type != "access" {
			w.WriteHeader(401)
			w.Write([]byte("4invalid token"))
			return
		}
		ctx := context.WithValue(r.Context(), "user", infoFromToken.UserInfo)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
