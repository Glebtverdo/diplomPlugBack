package routeshandler

import (
	"context"
	"diplomPlugService/internal/models"
	"diplomPlugService/internal/services"
	plugBack "diplomPlugService/internal/transport/grpc/gen/plugBack"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *serverAPI) GetRequests(
	ctx context.Context,
	in *plugBack.EmptyRequest,
) (*plugBack.Requests, error) {
	user := ctx.Value(models.UserKeyForContext).(models.UserInfo)
	requests, err := services.GetAllUsersRequests(user.Id)
	if err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}
	requestList := []*plugBack.Request{}
	for _, req := range requests {
		var arr []float32
		arr = append(arr, req.Object.Coords[0])
		arr = append(arr, req.Object.Coords[1])
		requestList = append(requestList, &plugBack.Request{
			Id: int32(req.Id),
			Object: &plugBack.Object{
				Name:    req.Object.Name,
				Address: req.Object.Address,
				Coords:  arr,
			},
		})
	}
	return &plugBack.Requests{Data: requestList}, nil
}
