package services

import (
	"diplomPlugService/internal/database"
	"diplomPlugService/internal/models"
)

func GetAllObjs() ([]models.Object, error) {
	return database.GetAllObjs()
}

func CreateNewObject(obj models.ObjectBody) error {
	return database.CreateNewObject(obj)
}

func DeleteObject(id int) error {
	return database.DeleteObject(id)
}

func ChangeObj(obj models.Object) error {
	return database.ChangeObj(obj)
}
