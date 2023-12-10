package services

import (
	"diplomPlugService/internal/database"
	"diplomPlugService/internal/models"
)

func GetAllRequests() ([]models.Request, error) {
	return database.GetAllRequests()
}

func GetAllUsersRequests(userId int) ([]models.RequestNoEngeeners, error) {
	return database.GetAllUsersRequests(userId)
}

func CreateNewRequest(obj models.RequestBody) error {
	return database.CreateNewRequest(obj)
}

func DeleteRequest(id int) error {
	return database.DeleteRequest(id)
}

func ChangeRequest(obj models.Request) error {
	return database.ChangeRequest(obj)
}
