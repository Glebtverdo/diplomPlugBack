package models

type Request struct {
	Id        int            `json:"id"`
	Object    ObjectBody     `json:"object"`
	Engeeners []UserInfoNoId `json:"users"`
}

type RequestBody struct {
	ObjectId int   `json:"objectId"`
	UsersIds []int `json:"usersIds"`
}
