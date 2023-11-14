package services

import (
	"diplomPlugService/internal/database"
	"diplomPlugService/internal/models"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt"
)

func GetAllUsers() ([]models.User, error) {
	return database.GetAllUsers()
}

func CreateNewUser(user models.UserBody) error {
	fmt.Println(user.FirstName, user.LastName, user.Login, user.MiddleName)
	return database.CreateNewUser(user)
}

func DeleteUser(id int) error {
	return database.DeleteUser(id)
}

func ChangeUser(user models.User) error {
	return database.ChangeUser(user)
}

func Login(loginPair models.UserLoginStruct) (models.JwtTokenPair, error) {
	user := database.GetUserByLoginAndPassword(loginPair)
	payload := models.JwtClaims{
		Type: "access",
		User: user,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 72).Unix(),
			Issuer:    "test",
		},
	}
	var (
		pair models.JwtTokenPair
		e    error
	)
	// Создаем новый JWT-токен и подписываем его по алгоритму HS256
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	pair.Access, e = token.SignedString([]byte("secret word"))
	if e != nil {
		return pair, e
	}
	payload.Type = "refresh"
	token = jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	pair.Refresh, e = token.SignedString([]byte("secret word"))
	if e != nil {
		return pair, e
	}
	return pair, e
}

func Logout() error {
	return nil
}
