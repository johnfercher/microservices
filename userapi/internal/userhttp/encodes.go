package userhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/go-kit/kit/endpoint"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"github.com/johnfercher/microservices/userapi/pkg/api/apifields"
	"github.com/johnfercher/microservices/userapi/pkg/api/apilog"
	"net/http"
)

const (
	UnknownError string = "unknown_error"
)

func EncodeResponse(ctx context.Context, w http.ResponseWriter, response interface{}) error {
	w.Header().Set("Content-Type", "application/json")
	LogResponse(ctx, response)
	return json.NewEncoder(w).Encode(response)
}

func EncodeError(ctx context.Context, err error, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")

	if apiError, ok := err.(apierror.ApiError); ok {
		w.WriteHeader(apiError.GetStatusCode())
		_ = json.NewEncoder(w).Encode(apiError)
		LogResponse(ctx, apiError)
		return
	}

	httpStatus := http.StatusInternalServerError

	unknownError := apierror.New(ctx, UnknownError, httpStatus).
		AppendFields(apifields.String("err", err.Error()))

	apierror.Log(ctx, unknownError)

	w.WriteHeader(httpStatus)
	_ = json.NewEncoder(w).Encode(unknownError)
	LogResponse(ctx, unknownError)

	return
}

func LogResponse(ctx context.Context, response interface{}) {
	responseBytes, err := json.Marshal(response)
	if err != nil {
		apilog.Info(ctx, fmt.Sprintf("Cannot encode response: %v", err))
		return
	}

	if len(responseBytes) > 5000 {
		apilog.Info(ctx, "Response larger than 5Kb")
	}

	apilog.Info(ctx, fmt.Sprintf("Response returned: %s", responseBytes))
}

func LogRequest(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		requestBytes, erMarshal := json.Marshal(request)
		if erMarshal != nil {
			apilog.Info(ctx, fmt.Sprintf("Cannot encode request: %v", erMarshal))
		}

		apilog.Info(ctx, fmt.Sprintf("Request received: %s", requestBytes))
		return next(ctx, request)
	}
}
