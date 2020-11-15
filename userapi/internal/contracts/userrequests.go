package contracts

import (
	"github.com/google/uuid"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
)

type CreateUserRequest struct {
	Name string `json:"name"`
	Type int    `json:"type"`
}

func (self *CreateUserRequest) ToUser() (*entity.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &entity.User{
		Id:     id.String(),
		Name:   self.Name,
		Type:   self.Type,
		Active: true,
	}, nil
}

type UpdateUserRequest struct {
	Id   string `json:"-"`
	Name string `json:"name"`
	Type int    `json:"type"`
}

func (self *UpdateUserRequest) ToUser() *entity.User {
	return &entity.User{
		Id:     self.Id,
		Name:   self.Name,
		Type:   self.Type,
		Active: true,
	}
}
