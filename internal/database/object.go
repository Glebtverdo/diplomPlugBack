package database

import (
	"context"
	"diplomPlugService/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetAllObjs() ([]models.Object, error) {
	rows, err := pool.Query(context.Background(), "Select * from objects order by id")
	if err != nil {
		return nil, err
	}
	var arr []models.Object
	for rows.Next() {
		var obj models.Object
		rows.Scan(&obj.Id, &obj.Name, &obj.Address, &obj.Coords[0], &obj.Coords[1])
		arr = append(arr, obj)
	}
	defer rows.Close()
	return arr, nil
}

func CreateNewObject(obj models.ObjectBody) error {
	query := ("Insert into objects (name, address, firstCoord, secondCoord) values(@objectName, @objectAddress, @firstCoord, @secondCoord) RETURNING *")
	args := pgx.NamedArgs{
		"objectName":    obj.Name,
		"objectAddress": obj.Address,
		"firstCoord":    obj.Coords[0],
		"secondCoord":   obj.Coords[1],
	}
	res, err := pool.Exec(context.Background(), query, args)
	if err != nil {
		return err
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("can not create this object")
	}
	return nil
}

func DeleteObject(id int) error {
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

func ChangeObj(obj models.Object) error {
	query := ("Update objects SET (name, address firstCoord, secondCoord) = (@objectName, @objectAddress, @firstCoord, @secondCoord) Where id = @id")
	args := pgx.NamedArgs{
		"id":            obj.Id,
		"objectName":    obj.Name,
		"objectAddress": obj.Address,
		"firstCoord":    obj.Coords[0],
		"secondCoord":   obj.Coords[1],
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
