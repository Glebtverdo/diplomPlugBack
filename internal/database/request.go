package database

import (
	"context"
	"diplomPlugService/internal/models"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func GetAllRequests() ([]models.Request, error) {
	rows, err := pool.Query(context.Background(),
		`Select requests.id as id, name, address, firstCoord,
			secondCoord, firstname, lastname, middlename  from requests 
			Join objects on requests.objectId = objects.id
			Join users_request on requests.id = users_request.requestId
			Join users on users.id = users_request.userId
			`)
	if err != nil {
		return nil, err
	}
	var arr []models.Request
	writtenRequests := make(map[int]models.Request)
	for rows.Next() {
		var request models.Request
		var user models.UserInfoNoId
		rows.Scan(&request.Id, &request.Object.Name, &request.Object.Address,
			&request.Object.Coords[0], &request.Object.Coords[1],
			&user.FirstName, &user.LastName, &user.MiddleName)
		_, ok := writtenRequests[request.Id]
		if !ok {
			writtenRequests[request.Id] = request
		}
		request.Engeeners = append(writtenRequests[request.Id].Engeeners, user)
		writtenRequests[request.Id] = request
		arr = append(arr, request)
	}
	defer rows.Close()
	return arr, nil
}

func CreateNewRequest(request models.RequestBody) error {
	query := ("Insert into requests (objectId) values($1) RETURNING id")
	// args := pgx.NamedArgs{
	// 	"objectName": request.ObjectId,
	// }
	ctx := context.Background()
	transaction, err := pool.Begin(ctx)
	if err != nil {
		return err
	}

	row := transaction.QueryRow(ctx, query, request.ObjectId)
	var requestId int
	row.Scan(&requestId)
	if requestId == 0 {
		transaction.Rollback(ctx)
		return err
	}
	query = ("Insert into users_request (requestId, userId) values(@requestId, @userId)")
	for _, userId := range request.UsersIds {
		args := pgx.NamedArgs{
			"requestId": requestId,
			"userId":    userId,
		}
		res, err := transaction.Exec(ctx, query, args)
		if err != nil || res.RowsAffected() == 0 {
			transaction.Rollback(ctx)
			return fmt.Errorf("error with user %d", userId)
		}
	}
	err = transaction.Commit(ctx)
	if err != nil {
		return err
	}
	return nil
}

func DeleteRequest(id int) error {
	query := ("Delete from users_request Where requestId = @id")
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
	query = ("Delete from requests Where id = @id")
	res, queryErr = pool.Exec(context.Background(), query, args)
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

func GetAllUsersRequests(userId int) ([]models.RequestNoEngeeners, error) {
	rows, err := pool.Query(context.Background(),
		`Select requests.id as id, name, address, firstCoord,
			secondCoord from requests 
			Join objects on requests.objectId = objects.id
			Join users_request on requests.id = users_request.requestId
			Where users_request.userId = $1
		`, userId)
	if err != nil {
		return nil, err
	}
	var requests []models.RequestNoEngeeners
	for rows.Next() {
		var request models.RequestNoEngeeners
		rows.Scan(&request.Id, &request.Object.Name, &request.Object.Address,
			&request.Object.Coords[0], &request.Object.Coords[1])
		requests = append(requests, request)
	}
	defer rows.Close()
	return requests, nil
}
