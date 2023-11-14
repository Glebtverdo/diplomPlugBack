package models

import (
	"github.com/golang-jwt/jwt"
)

type User struct {
	Id         int    `json:"id"`
	Login      string `json:"login"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
}

type UserBody struct {
	Login      string `json:"login"`
	Password   string `json:"password"`
	FirstName  string `json:"firstName"`
	LastName   string `json:"lastName"`
	MiddleName string `json:"middleName"`
}

type UserLoginStruct struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}

type JwtClaims struct {
	Type string
	User
	jwt.StandardClaims
}

type JwtTokenPair struct {
	Access  string `json:"access"`
	Refresh string `json:"refresh"`
}
