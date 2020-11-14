package userservice

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/internal/domain/repository"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
)

type userService struct {
	userRepository repository.UserRepository
}

func NewUserService(userRepository repository.UserRepository) *userService {
	return &userService{
		userRepository: userRepository,
	}
}

func (self *userService) Create(ctx context.Context, createRequest *contracts.CreateUserRequest) (*entity.User, apierror.ApiError) {
	user, err := createRequest.ToUser(ctx)
	if err != nil {
		return nil, err
	}

	err = self.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userService) GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	return self.userRepository.GetById(ctx, id)
}
