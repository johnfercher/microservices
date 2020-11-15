package userrepository

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
)

const (
	cannotExecuteQueryError string = "cannot_execute_query_error"
)

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (self *userRepository) Create(ctx context.Context, user *entity.User) apierror.ApiError {
	tx := self.db.Create(user)

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	return nil
}

func (self *userRepository) GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	user := &entity.User{}
	tx := self.db.Where("id = ?", id).First(user)

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	return user, nil
}

func (self *userRepository) Update(ctx context.Context, user *entity.User) apierror.ApiError {
	tx := self.db.Model(user).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"name": user.Name,
		"type": user.Type,
	})

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	return nil
}

func (self *userRepository) Deactivate(ctx context.Context, id string) apierror.ApiError {
	tx := self.db.Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"active": false,
	})

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	return nil
}

func (self *userRepository) Activate(ctx context.Context, id string) apierror.ApiError {
	tx := self.db.Model(&entity.User{}).Where("id = ?", id).Updates(map[string]interface{}{
		"active": true,
	})

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	return nil
}
