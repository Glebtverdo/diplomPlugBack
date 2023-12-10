package routeshandler

import (
	// Сгенерированный код
	"context"
	"diplomPlugService/internal/models"
	"diplomPlugService/internal/services"
	plugBack "diplomPlugService/internal/transport/grpc/gen/plugBack"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) Login(
	ctx context.Context,
	in *plugBack.LoginRequest,
) (*plugBack.LoginResponse, error) {
	loginPair := models.UserLoginStruct{
		Login:    in.Login,
		Password: in.Password,
	}
	jwtPair, err := services.Login(loginPair)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	return &plugBack.LoginResponse{Access: jwtPair.Access, Refresh: jwtPair.Refresh}, nil
}
