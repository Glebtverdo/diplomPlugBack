package services

import (
	"diplomPlugService/internal/models"

	"github.com/golang-jwt/jwt/v5"
)

func IsAuthorizate(unParsingToken string) {
	println(unParsingToken)
	var claims jwt.Claims
	res, _ := jwt.ParseWithClaims(unParsingToken, claims, func(token *jwt.Token) (interface{}, error) {
		println(token)
		var some models.JwtClaims
		return some, nil
	})
	println(res)
}
