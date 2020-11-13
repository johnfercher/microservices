package repository

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
)

type UserRepository interface {
	GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError)
}
