package models

type Request struct {
	Id     int    `json:"id"`
	Object Object `json:"object"`
	Users  []User `json:"users"`
}

type RequestBody struct {
	ObjectId int   `json:"objectId"`
	UsersIds []int `json:"usersIds"`
}
