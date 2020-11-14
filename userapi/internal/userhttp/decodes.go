package userhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"go.uber.org/zap"
	"io/ioutil"
	"net/http"
)

const (
	InvalidUrlParametersError = "invalid_url_parameter_error"
	EmptyBodyError            = "empty_body_error"
	DecodebodyError           = "decode_body_error"
)

func DecodeIdFromUrl(ctx context.Context, r *http.Request) (interface{}, error) {
	uriParams := mux.Vars(r)

	id := uriParams["id"]
	if id == "" {
		err := apierror.New(ctx, InvalidUrlParametersError, http.StatusBadRequest).
			WithMessage(fmt.Sprint("Id cannot be empty"))

		apierror.Log(ctx, err)
		return nil, err
	}

	return id, nil
}

func DecodeCreateUserRequestFromBody(ctx context.Context, r *http.Request) (interface{}, error) {
	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := apierror.New(ctx, EmptyBodyError, http.StatusBadRequest).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	createRequest := &contracts.CreateUserRequest{}

	err = json.Unmarshal(bytes, createRequest)
	if err != nil {
		apiErr := apierror.New(ctx, DecodebodyError, http.StatusBadRequest).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	return createRequest, nil
}
