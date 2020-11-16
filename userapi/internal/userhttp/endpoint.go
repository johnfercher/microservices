package userhttp

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/internal/domain/service"
)

func MakeGetByIdEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		return userService.GetById(ctx, id)
	}
}

func MakeSearchEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		searchRequest := request.(*contracts.SearchRequest)
		return userService.Search(ctx, searchRequest)
	}
}

func MakeCreateEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		createUserRequest := request.(*contracts.CreateUserRequest)
		return userService.Create(ctx, createUserRequest)
	}
}

func MakeUpdateEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		updateUserRequest := request.(*contracts.UpdateUserRequest)
		return userService.Update(ctx, updateUserRequest)
	}
}

func MakeDeactivateEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		return userService.Deactivate(ctx, id)
	}
}

func MakeActivateEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		id := request.(string)
		return userService.Activate(ctx, id)
	}
}

func MakeAddTypeEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		userType := request.(*entity.Type)
		return userService.AddUserType(ctx, userType)
	}
}

func MakeRemoveTypeEndpoint(userService service.UserService) func(ctx context.Context, request interface{}) (interface{}, error) {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		userType := request.(*entity.Type)
		return userService.RemoveUserType(ctx, userType)
	}
}
