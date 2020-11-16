package contracts

import (
	"github.com/google/uuid"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
)

type CreateUserRequest struct {
	Name  string        `json:"name"`
	Types []entity.Type `json:"types"`
}

func (self *CreateUserRequest) ToUser() (*entity.User, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	var types []entity.Type
	for _, userType := range self.Types {
		types = append(types, userTypeStringToEntity(userType.Type, id.String()))
	}

	return &entity.User{
		Id:     id.String(),
		Name:   self.Name,
		Types:  types,
		Active: true,
	}, nil
}

type UpdateUserRequest struct {
	Id   string `json:"-"`
	Name string `json:"name"`
}

func (self *UpdateUserRequest) ToUser() *entity.User {
	return &entity.User{
		Id:     self.Id,
		Name:   self.Name,
		Active: true,
	}
}

func userTypeStringToEntity(userType string, userId string) entity.Type {
	return entity.Type{
		UserId: userId,
		Type:   userType,
	}
}

type SearchRequest struct {
	Id     *string
	Name   *string
	Active *bool
	Type   *string
	Limit  int64
	Offset int64
}
