package service

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
)

type UserService interface {
	Create(ctx context.Context, createRequest *contracts.CreateUserRequest) (*entity.User, apierror.ApiError)
	GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError)
}
