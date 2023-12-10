package grpcTrasport

import (
	routeshandler "diplomPlugService/internal/transport/grpc/routesHandler"
	"diplomPlugService/internal/transport/midlewares"
	"log"
	"net"

	"google.golang.org/grpc"
)

// type App struct {
// 	gRPCServer *grpc.Server
// 	port       int // Порт, на котором будет работать grpc-сервер
// }

// func InitGrpcServer(loginService routeshandler.Login, port int) *App {
// 	gRPCServer := grpc.NewServer()
// 	routeshandler.Register(gRPCServer, loginService)
// 	return &App{
// 		gRPCServer: gRPCServer,
// 		port:       2000,
// 	}
// }

// func (a *App) Run() error {
// 	const op = "grpcapp.Run"

// 	// Создаём listener, который будет слушить TCP-сообщения, адресованные
// 	// Нашему gRPC-серверу
// 	l, err := net.Listen("tcp", fmt.Sprintf(":%d", a.port))
// 	if err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

// 	// Запускаем обработчик gRPC-сообщений
// 	if err := a.gRPCServer.Serve(l); err != nil {
// 		return fmt.Errorf("%s: %w", op, err)
// 	}

//		return nil
//	}

func InitServer() {
	lis, err := net.Listen("tcp", ":2000")
	if err != nil {
		log.Fatalf("failed to listen grpc: %v", err)
	}
	s := grpc.NewServer(
		grpc.ChainStreamInterceptor(midlewares.CheckAuthorizationStreamInterceptor),
		grpc.ChainUnaryInterceptor(midlewares.CheckAuthorizationUnaryInterceptor),
	)
	routeshandler.Register(s)
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
