package userhttp

import (
	"context"
	"encoding/json"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"go.uber.org/zap"
	"net/http"
)

const (
	UnknownError string = "unknown_error"
)

func EncodeResponse(_ context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if apiError, ok := err.(apierror.ApiError); ok {
		w.WriteHeader(apiError.GetStatusCode())
		_ = json.NewEncoder(w).Encode(apiError)
		return
	}

	httpStatus := http.StatusInternalServerError

	unknownError := apierror.New(ctx, UnknownError, httpStatus).
		AppendFields(zap.String("err", err.Error()))

	apierror.Log(ctx, unknownError)

	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(unknownError)
	return
}
