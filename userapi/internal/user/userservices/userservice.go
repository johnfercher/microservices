package userservices

import (
	"context"
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

func (self *userService) Create(ctx context.Context, user *entity.User) apierror.ApiError {
	return nil
}

func (self *userService) GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	user, err := self.userRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
