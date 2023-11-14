package database

import (
	"context"
	"diplomPlugService/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetAllRequests() ([]models.Request, error) {
	rows, err := pool.Query(context.Background(), "Select * from requests")
	if err != nil {
		return nil, err
	}
	var arr []models.Request
	for rows.Next() {
		var obj models.Request
		rows.Scan(&obj.Id)
		arr = append(arr, obj)
	}
	defer rows.Close()
	return arr, nil
}

func CreateNewRequest(request models.RequestBody) error {
	query := ("Insert into objects (name, address) values(@objectName, @objectAddress) RETURNING *")
	args := pgx.NamedArgs{
		"objectName":    request.UsersIds,
		"objectAddress": request.ObjectId,
	}
	res, err := pool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("can not create this object")
	}
	// row := pool.QueryRow(context.Background(), "Select * from  objects where name = @objectName", args)
	// var resultObj models.Request
	// row.Scan(&resultObj.Id, &resultObj.Name, &resultObj.Address)
	return nil
}

func DeleteRequest(id int) error {
	query := ("Delete from objects Where id = @id")
	args := pgx.NamedArgs{
		"id": id,
	}
	res, queryErr := pool.Exec(context.Background(), query, args)
	if queryErr != nil {
		return queryErr
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("object with id = %d, does not exists", id)
	}
	return nil
}

func ChangeRequest(obj models.Request) error {
	query := ("Update objects SET (name, address) = (@objectName, @objectAddress) Where id = @id")
	args := pgx.NamedArgs{
		"id": obj.Id,
		// "objectName":    obj.Name,
		// "objectAddress": obj.Address,
	}
	res, queryErr := pool.Exec(context.Background(), query, args)
	if queryErr != nil {
		return queryErr
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("object with id = %d, does not exists", obj.Id)
	}
	return nil
}
