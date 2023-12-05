package models

type Object struct {
	Id      int        `json:"id"`
	Name    string     `json:"name"`
	Address string     `json:"address"`
	Coords  [2]float32 `json:"coords"`
}

type ObjectBody struct {
	Name    string     `json:"name"`
	Address string     `json:"address"`
	Coords  [2]float32 `json:"coords"`
}
