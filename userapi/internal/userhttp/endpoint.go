package userhttp

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/service"
)

func MakeGetByIdEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		return userService.GetById(ctx, id)
	}
}

func MakeCreateEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		createUserRequest := request.(*contracts.CreateUserRequest)
		return userService.Create(ctx, createUserRequest)
	}
}
