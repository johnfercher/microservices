package userservice

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/internal/domain/service"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
)

type userEvents struct {
	inner     service.UserService
	publisher service.TopicPublisher
}

func NewUserEvents(inner service.UserService, publisher service.TopicPublisher) *userEvents {
	return &userEvents{
		inner:     inner,
		publisher: publisher,
	}
}

func (self *userEvents) Create(ctx context.Context, createRequest *contracts.CreateUserRequest) (*entity.User, apierror.ApiError) {
	user, err := self.inner.Create(ctx, createRequest)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userEvents) GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	user, err := self.inner.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userEvents) Update(ctx context.Context, updateRequest *contracts.UpdateUserRequest) (*entity.User, apierror.ApiError) {
	user, err := self.inner.Update(ctx, updateRequest)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userEvents) Deactivate(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	user, err := self.inner.Deactivate(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (self *userEvents) Activate(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	user, err := self.inner.Activate(ctx, id)
	if err != nil {
		return nil, err
	}

	return user, nil
}
