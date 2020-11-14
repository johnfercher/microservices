package repository

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
)

type UserRepository interface {
	Create(ctx context.Context, user *entity.User) apierror.ApiError
	GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError)
}
