package userservice

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/internal/domain/repository"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"go.uber.org/zap"
	"net/http"
)

const (
	cannotCreateUuidError string = "cannot_create_uuid_error"
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
	user, errUserCreation := createRequest.ToUser()
	if errUserCreation != nil {
		apiError := apierror.New(ctx, cannotCreateUuidError, http.StatusInternalServerError).
			AppendFields(zap.String("err", errUserCreation.Error()))

		apierror.Log(ctx, apiError)
		return nil, apiError
	}

	err := self.userRepository.Create(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userService) GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	return self.userRepository.GetById(ctx, id)
}

func (self *userService) Update(ctx context.Context, updateRequest *contracts.UpdateUserRequest) (*entity.User, apierror.ApiError) {
	user := updateRequest.ToUser()

	err := self.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userService) Deactivate(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	err := self.userRepository.Deactivate(ctx, id)
	if err != nil {
		return nil, err
	}

	return self.userRepository.GetById(ctx, id)
}

func (self *userService) Activate(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	err := self.userRepository.Activate(ctx, id)
	if err != nil {
		return nil, err
	}

	return self.userRepository.GetById(ctx, id)
}
