package repository

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
)

type UserRepository interface {
	// Read
	GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError)
	Search(ctx context.Context, searchRequest *contracts.SearchRequest) (*entity.Page, apierror.ApiError)

	// Write
	Create(ctx context.Context, user *entity.User) apierror.ApiError
	AddUserType(ctx context.Context, userType *entity.Type) apierror.ApiError
	RemoveUserType(ctx context.Context, userType *entity.Type) apierror.ApiError
	Update(ctx context.Context, user *entity.User) apierror.ApiError
	Deactivate(ctx context.Context, id string) apierror.ApiError
	Activate(ctx context.Context, id string) apierror.ApiError
}
