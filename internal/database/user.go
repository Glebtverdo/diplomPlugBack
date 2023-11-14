package database

import (
	"context"
	"diplomPlugService/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetAllUsers() ([]models.User, error) {
	rows, err := pool.Query(context.Background(), "Select * from users")
	if err != nil {
		return nil, err
	}
	var arr []models.User
	for rows.Next() {
		var userInfo models.User
		rows.Scan(&userInfo.Id, &userInfo.Login, &userInfo.Password, &userInfo.FirstName, &userInfo.LastName, &userInfo.MiddleName)
		arr = append(arr, userInfo)
	}
	defer rows.Close()
	return arr, nil
}

func CreateNewUser(user models.UserBody) error {
	query := (`Insert into users (login, password, firstName, lastName, middleName)
		 values(@login, @password, @firstName, @lastName, @middleName)`)
	args := pgx.NamedArgs{
		"login":      user.Login,
		"password":   user.Password,
		"firstName":  user.FirstName,
		"lastName":   user.LastName,
		"middleName": user.MiddleName,
	}
	res, err := pool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("can not create this user")
	}
	return nil
}

func DeleteUser(id int) error {
	return nil
}

func ChangeUser(user models.User) error {
	query := (`Update objects SET (login, password, firstName, lastName, middleName) = 
	(@login, @password, @firstName, @lastName, @middleName)`)
	args := pgx.NamedArgs{
		"login":      user.Login,
		"password":   user.Password,
		"firstName":  user.FirstName,
		"lastName":   user.LastName,
		"middleName": user.MiddleName,
	}
	res, queryErr := pool.Exec(context.Background(), query, args)
	if queryErr != nil {
		return queryErr
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("user does not exists")
	}
	return nil
}

func GetUserByLoginAndPassword(user models.UserLoginStruct) models.User {
	query := fmt.Sprintf("Select * from users Where login = %s and password = %s", user.Login, user.Password)
	row := pool.QueryRow(context.Background(), query)
	var userInfo models.User
	row.Scan(&userInfo.Id, &userInfo.Login, &userInfo.Password, &userInfo.FirstName, &userInfo.LastName, &userInfo.MiddleName)
	return userInfo
}
