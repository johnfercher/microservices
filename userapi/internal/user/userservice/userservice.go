package userservice

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/internal/domain/repository"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"github.com/johnfercher/microservices/userapi/pkg/api/apifields"
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
			AppendFields(apifields.String("err", errUserCreation.Error()))

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

func (self *userService) Search(ctx context.Context, searchRequest *contracts.SearchRequest) (*entity.Page, apierror.ApiError) {
	return self.userRepository.Search(ctx, searchRequest)
}

func (self *userService) Update(ctx context.Context, updateRequest *contracts.UpdateUserRequest) (*entity.User, apierror.ApiError) {
	user := updateRequest.ToUser()

	err := self.userRepository.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return self.GetById(ctx, updateRequest.Id)
}

func (self *userService) AddUserType(ctx context.Context, userType *entity.Type) (*entity.User, apierror.ApiError) {
	err := self.userRepository.AddUserType(ctx, userType)
	if err != nil {
		return nil, err
	}

	return self.GetById(ctx, userType.UserId)
}

func (self *userService) RemoveUserType(ctx context.Context, userType *entity.Type) (*entity.User, apierror.ApiError) {
	err := self.userRepository.RemoveUserType(ctx, userType)
	if err != nil {
		return nil, err
	}

	return self.GetById(ctx, userType.UserId)
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
