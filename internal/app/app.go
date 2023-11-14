package app

import (
	"diplomPlugService/internal/transport/rest"
)

func RunProject() {
	// dbError := database.InitDatabase()
	// if dbError != nil {
	// 	fmt.Println(dbError)
	// 	return
	// }
	rest.InitServer()
}
