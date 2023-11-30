package app

import (
	"diplomPlugService/internal/database"
	"diplomPlugService/internal/transport/rest"
	"fmt"
)

func RunProject() {
	dbError := database.InitDatabase()
	if dbError != nil {
		fmt.Println(dbError)
		return
	}
	rest.InitServer()
}
