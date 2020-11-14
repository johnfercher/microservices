package contracts

import (
	"context"
	"github.com/google/uuid"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"net/http"
)

const (
	cannotCreateUuidError string = "cannot_create_uuid_error"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Type int    `json:"type,omitempty"`
}

func (self *CreateUserRequest) ToUser(ctx context.Context) (*entity.User, apierror.ApiError) {
	id, err := uuid.NewRandom()
	if err != nil {
		apiError := apierror.New(ctx, cannotCreateUuidError, http.StatusInternalServerError)

		apierror.Log(ctx, apiError)
		return nil, apiError
	}

	return &entity.User{
		Id:   id.String(),
		Name: self.Name,
		Type: self.Type,
	}, nil
}
