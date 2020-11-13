package userhttp

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"net/http"
)

const (
	InvalidUrlParametersError = "invalid_url_parameter_error"
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
