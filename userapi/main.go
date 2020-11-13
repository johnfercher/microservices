package main

import (
	"github.com/gorilla/mux"
	"github.com/johnfercher/microservices/userapi/internal/user/userservices"
	"github.com/johnfercher/microservices/userapi/internal/userhttp"
	"github.com/johnfercher/microservices/userapi/pkg/api/apiscope"
	"log"
	"net/http"

	httptransport "github.com/go-kit/kit/transport/http"
)

func main() {
	userService := userservices.NewUserService()

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(userhttp.EncodeError),
	}

	getUserById := httptransport.NewServer(
		userhttp.MakeGetByIdEndpoint(userService),
		userhttp.DecodeIdFromUrl,
		userhttp.EncodeResponse,
		serverOptions...,
	)

	router := mux.NewRouter()
	router.Use(apiscope.LifecycleCtxSetup())

	router.Handle("/uppercase/{id}", getUserById).Methods(http.MethodGet)

	log.Fatal(http.ListenAndServe(":8080", router))
}
