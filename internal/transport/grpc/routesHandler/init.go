package routeshandler

import (
	plugBack "diplomPlugService/internal/transport/grpc/gen/plugBack"

	"google.golang.org/grpc"
)

type serverAPI struct {
	plugBack.UnimplementedPlugBackGrpcServer
}

func Register(gRPCServer *grpc.Server) {
	plugBack.RegisterPlugBackGrpcServer(gRPCServer, &serverAPI{})
}
