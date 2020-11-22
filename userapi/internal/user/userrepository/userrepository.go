package userrepository

import (
	"context"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"go.uber.org/zap"
	"gorm.io/gorm"
	"net/http"
	"strings"
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
	err := self.db.Transaction(func(tx *gorm.DB) error {
		if tx.Create(user).Error != nil {
			apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
				WithMessage("It was not possible create user").
				AppendFields(zap.String("err", tx.Error.Error()))

			apierror.Log(ctx, apiErr)
			return apiErr
		}

		return nil
	})

	if err != nil {
		return err.(apierror.ApiError)
	}

	return nil
}

func (self *userRepository) AddUserType(ctx context.Context, userType *entity.Type) apierror.ApiError {
	tx := self.db.Create(userType)

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	return nil
}

func (self *userRepository) RemoveUserType(ctx context.Context, userType *entity.Type) apierror.ApiError {
	tx := self.db.Where("user_id = ? AND type = ?", userType.UserId, userType.Type).Delete(userType)

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
			WithMessage("Cannot load user").
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	err := self.db.Model(user).Association("Types").Find(&user.Types)
	if err != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			WithMessage("Cannot load types from user").
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	return user, nil
}

func (self *userRepository) Search(ctx context.Context, searchRequest *contracts.SearchRequest) (*entity.Page, apierror.ApiError) {
	page := &entity.Page{
		Paging: entity.Paging{
			Offset: searchRequest.Offset,
			Limit:  searchRequest.Limit,
		},
		Results: []entity.User{},
	}

	query := []string{}
	args := []interface{}{}

	if searchRequest.Id != nil {
		query = append(query, "users.id = ?")
		args = append(args, *searchRequest.Id)
	}

	if searchRequest.Name != nil {
		query = append(query, "users.name = ?")
		args = append(args, *searchRequest.Name)
	}

	if searchRequest.Active != nil {
		query = append(query, "users.active = ?")
		args = append(args, *searchRequest.Active)
	}

	if searchRequest.Type != nil {
		query = append(query, "types.type = ?")
		args = append(args, *searchRequest.Type)
	}

	tx := self.db.Table("users").
		Distinct("users.id").
		Select("users.id, users.name, users.active").
		Joins("left join types on types.user_id = users.id")

	if len(query) != 0 {
		tx = tx.Where(strings.Join(query, " AND "), args...)
	}

	tx = tx.Limit(int(searchRequest.Limit))

	tx = tx.Scan(&page.Results)

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			WithMessage("Search user").
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	var count int64 = 0

	tx.Count(&count)

	if tx.Error != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			WithMessage("Search count user").
			AppendFields(zap.String("err", tx.Error.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	page.Paging.Total = count

	return page, nil
}

func (self *userRepository) Update(ctx context.Context, user *entity.User) apierror.ApiError {
	tx := self.db.Model(&entity.User{}).Where("id = ?", user.Id).Updates(map[string]interface{}{
		"name": user.Name,
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
