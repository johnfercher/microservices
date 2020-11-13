package userrepository

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
)

type userRepository struct {
}

func NewUserRepository() *userRepository {
	return &userRepository{}
}

func (self *userRepository) Create(ctx context.Context, user *entity.User) apierror.ApiError {
	return nil
}

func (self *userRepository) GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	return &entity.User{
		Id:   "user_id",
		Name: "user_name",
		Address: &entity.Address{
			Code: "address",
		},
		Login:    "login",
		Password: "password",
	}, nil
}
