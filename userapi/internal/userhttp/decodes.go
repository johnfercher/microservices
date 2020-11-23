package userhttp

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnfercher/microservices/userapi/internal/contracts"
	"github.com/johnfercher/microservices/userapi/internal/domain/entity"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"github.com/johnfercher/microservices/userapi/pkg/api/apifields"
	"io/ioutil"
	"net/http"
	"strconv"
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
			AppendFields(apifields.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	createRequest := &contracts.CreateUserRequest{}

	err = json.Unmarshal(bytes, createRequest)
	if err != nil {
		apiErr := apierror.New(ctx, DecodebodyError, http.StatusBadRequest).
			AppendFields(apifields.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	return createRequest, nil
}

func DecodeUpdateUserRequestFromUrlAndBody(ctx context.Context, r *http.Request) (interface{}, error) {
	uriParams := mux.Vars(r)

	id := uriParams["id"]
	if id == "" {
		err := apierror.New(ctx, InvalidUrlParametersError, http.StatusBadRequest).
			WithMessage(fmt.Sprint("Id cannot be empty"))

		apierror.Log(ctx, err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := apierror.New(ctx, EmptyBodyError, http.StatusBadRequest).
			AppendFields(apifields.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	updateRequest := &contracts.UpdateUserRequest{}

	err = json.Unmarshal(bytes, updateRequest)
	if err != nil {
		apiErr := apierror.New(ctx, DecodebodyError, http.StatusBadRequest).
			AppendFields(apifields.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	updateRequest.Id = id

	return updateRequest, nil
}

func DecodeUserTypeFromUrlAndBody(ctx context.Context, r *http.Request) (interface{}, error) {
	uriParams := mux.Vars(r)

	id := uriParams["id"]
	if id == "" {
		err := apierror.New(ctx, InvalidUrlParametersError, http.StatusBadRequest).
			WithMessage(fmt.Sprint("Id cannot be empty"))

		apierror.Log(ctx, err)
		return nil, err
	}

	bytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		apiErr := apierror.New(ctx, EmptyBodyError, http.StatusBadRequest).
			AppendFields(apifields.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	userType := &entity.Type{}

	err = json.Unmarshal(bytes, userType)
	if err != nil {
		apiErr := apierror.New(ctx, DecodebodyError, http.StatusBadRequest).
			AppendFields(apifields.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return nil, apiErr
	}

	userType.UserId = id

	return userType, nil
}

func DecodeSearchFromUrl(ctx context.Context, r *http.Request) (interface{}, error) {
	var limitMax int64 = 100
	searchRequest := &contracts.SearchRequest{
		Limit:  30,
		Offset: 0,
	}

	queryParams := r.URL.Query()

	id := queryParams.Get("id")
	if id != "" {
		searchRequest.Id = &id
	}

	name := queryParams.Get("name")
	if name != "" {
		searchRequest.Name = &name
	}

	userType := queryParams.Get("type")
	if userType != "" {
		searchRequest.Type = &userType
	}

	activeString := queryParams.Get("active")
	if activeString != "" {
		active, err := strconv.ParseBool(activeString)
		if err == nil {
			searchRequest.Active = &active
		}
	}

	limitString := queryParams.Get("limit")
	if limitString != "" {
		limit, err := strconv.ParseInt(limitString, 10, 64)
		if err == nil {
			if limit > limitMax {
				limit = limitMax
			}

			searchRequest.Limit = limit
		}
	}

	offsetString := queryParams.Get("offset")
	if offsetString != "" {
		offset, err := strconv.ParseInt(offsetString, 10, 64)
		if err == nil {
			searchRequest.Offset = offset
		}
	}

	return searchRequest, nil
}
