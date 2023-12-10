package app

import (
	"diplomPlugService/internal/database"
	grpcTrasport "diplomPlugService/internal/transport/grpc"
	restTransport "diplomPlugService/internal/transport/rest"
	"fmt"
	"sync"
)

func RunProject() {
	dbError := database.InitDatabase()
	if dbError != nil {
		fmt.Println(dbError)
		return
	}
	go restTransport.InitServer()
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		grpcTrasport.InitServer()
	}()

	wg.Wait()
}
