package userrepository

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"go.uber.org/zap"
	"net/http"
)

const (
	cannotExecuteQueryError  string = "cannot_execute_query_error"
	cannotFindUserErrorError string = "cannot_find_user_error"
	cannotParseObjectError   string = "cannot_parse_object_error"
)

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (self *userRepository) Create(ctx context.Context, user *entity.User) apierror.ApiError {
	_, err := self.db.ExecContext(ctx, fmt.Sprintf(`INSERT INTO main_user (id, name, user_type) values ('%s', '%s', '%d')`, user.Id, user.Name, user.Type))
	if err != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	return nil
}

func (self *userRepository) GetById(ctx context.Context, id string) (*entity.User, apierror.ApiError) {
	results, err := self.db.QueryContext(ctx, fmt.Sprintf(`SELECT id, name, user_type FROM main_user WHERE id = '%s';`, id))
	if err != nil {
		apiErr := apierror.New(ctx, cannotExecuteQueryError, http.StatusInternalServerError).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	var idRead sql.NullString
	var name sql.NullString
	var userType sql.NullInt32

	if !results.Next() {
		apiErr := apierror.New(ctx, cannotFindUserErrorError, http.StatusNotFound)

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	err = results.Scan(&idRead, &name, &userType)
	if err != nil {
		apiErr := apierror.New(ctx, cannotParseObjectError, http.StatusInternalServerError).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	user := entity.User{
		Id:   idRead.String,
		Name: name.String,
		Type: int(userType.Int32),
	}

	return &user, nil
}
