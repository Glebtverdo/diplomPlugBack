package models

type Object struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Address string `json:"address"`
}

type ObjectBody struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}
