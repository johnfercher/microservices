package userhttp

import (
	"context"
	"github.com/go-kit/kit/endpoint"
	"github.com/johnfercher/microservices/userapi/internal/domain/service"
)

func MakeGetByIdEndpoint(service service.UserService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		return service.GetById(ctx, id)
	}
}
