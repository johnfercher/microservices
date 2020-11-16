package main

import (
	"fmt"
	httptransport "github.com/go-kit/kit/transport/http"
	"github.com/gorilla/mux"
	"github.com/johnfercher/microservices/userapi/internal/infra"
	"github.com/johnfercher/microservices/userapi/internal/user/userrepository"
	"github.com/johnfercher/microservices/userapi/internal/user/userservice"
	"github.com/johnfercher/microservices/userapi/internal/userhttp"
	"github.com/johnfercher/microservices/userapi/pkg/api/apilog"
	"github.com/johnfercher/microservices/userapi/pkg/api/apiscope"
	"net/http"
)

var logger = apilog.New()

func main() {
	// Configs
	mysqlUrl := "localhost:3306"
	//mysqlUrl := "user-db:3306" // inside docker
	mysqlDbName := "UserDb"
	mysqlAdminUser := "AdminUser"
	mysqlAdminPassword := "AdminPassword"

	kafkaEventsUrl := "localhost:9092"
	kafkaEventsTopic := "topic-user-events"

	// Infra
	mysqlDb, err := infra.NewMysqlDb(mysqlUrl, mysqlDbName, mysqlAdminUser, mysqlAdminPassword)
	if err != nil {
		panic(err)
	}

	kafkaEventsPublisher := infra.NewTopicPublisher(kafkaEventsUrl, kafkaEventsTopic)

	// Repository
	userRepository := userrepository.NewUserRepository(mysqlDb)

	// Service
	userService := userservice.NewUserService(userRepository)
	userEvents := userservice.NewUserEvents(userService, kafkaEventsPublisher)

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(userhttp.EncodeError),
	}

	router := mux.NewRouter()
	router.Use(apiscope.LifecycleCtxSetup())

	RegisterEndpoint(router, "/users/search", http.MethodGet, httptransport.NewServer(
		userhttp.MakeSearchEndpoint(userEvents),
		userhttp.DecodeSearchFromUrl,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users/{id}", http.MethodGet, httptransport.NewServer(
		userhttp.MakeGetByIdEndpoint(userEvents),
		userhttp.DecodeIdFromUrl,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users", http.MethodPost, httptransport.NewServer(
		userhttp.MakeCreateEndpoint(userEvents),
		userhttp.DecodeCreateUserRequestFromBody,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users/{id}", http.MethodPut, httptransport.NewServer(
		userhttp.MakeUpdateEndpoint(userEvents),
		userhttp.DecodeUpdateUserRequestFromUrlAndBody,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users/{id}/active", http.MethodDelete, httptransport.NewServer(
		userhttp.MakeDeactivateEndpoint(userEvents),
		userhttp.DecodeIdFromUrl,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users/{id}/active", http.MethodPut, httptransport.NewServer(
		userhttp.MakeActivateEndpoint(userEvents),
		userhttp.DecodeIdFromUrl,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users/{id}/types", http.MethodPut, httptransport.NewServer(
		userhttp.MakeAddTypeEndpoint(userEvents),
		userhttp.DecodeUserTypeFromUrlAndBody,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users/{id}/types", http.MethodDelete, httptransport.NewServer(
		userhttp.MakeRemoveTypeEndpoint(userEvents),
		userhttp.DecodeUserTypeFromUrlAndBody,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Error(fmt.Sprintf("Shutdown %s", err.Error()))
	}
}

func RegisterEndpoint(router *mux.Router, path string, method string, server *httptransport.Server) {
	logger.Info(fmt.Sprintf("Registered -> Method:%s Path:%s", method, path))
	router.Handle(path, server).Methods(method)
}
