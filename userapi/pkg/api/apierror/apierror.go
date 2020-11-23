package apierror

import (
	"context"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/johnfercher/microservices/userapi/pkg/api/apifields"
	"github.com/johnfercher/microservices/userapi/pkg/api/apilog"
)

type ApiError interface {
	// Public
	GetStatusCode() int
	Error() string
	WithMessage(message string) ApiError
	AppendFields(fields ...apifields.Field) ApiError

	// Private
	getFields() []apifields.Field
}

func New(ctx context.Context, errorCode string, statusCode int) ApiError {
	id, _ := uuid.NewRandom()

	return &apiError{
		ErrorCode:  errorCode,
		StatusCode: statusCode,
		RequestId:  apilog.GetContextStringField(ctx, apilog.CtxRequestId),
		id:         id.String(),
	}
}

type apiError struct {
	// Public
	Message       string        `json:"message,omitempty"`
	ErrorCode     string        `json:"error_code,omitempty"`
	StatusCode    int           `json:"status_code,omitempty"`
	RequestId     string        `json:"request_id,omitempty"`
	Causes        []interface{} `json:"causes,omitempty"`
	RelatedErrors []apiError    `json:"related_errors,omitempty"`

	// Private
	id     string            `json:"-"`
	fields []apifields.Field `json:"-"`
}

func (self *apiError) WithMessage(message string) ApiError {
	self.Message = message
	return self
}

func (self *apiError) AppendFields(fields ...apifields.Field) ApiError {
	self.fields = append(self.fields, fields...)
	return self
}

func (self *apiError) GetStatusCode() int {
	return self.StatusCode
}

func (self *apiError) Error() string {
	bytesErr, _ := json.Marshal(self)
	return string(bytesErr)
}

func (self *apiError) getFields() []apifields.Field {
	return self.fields
}

func Log(ctx context.Context, err ApiError) {
	apilog.Error(ctx, err.Error(), err.getFields()...)
}
