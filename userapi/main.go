package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/johnfercher/microservices/userapi/internal/user/userrepository"
	"github.com/johnfercher/microservices/userapi/internal/user/userservices"
	"github.com/johnfercher/microservices/userapi/internal/userhttp"
	"github.com/johnfercher/microservices/userapi/pkg/api/apilog"
	"github.com/johnfercher/microservices/userapi/pkg/api/apiscope"
	"github.com/segmentio/kafka-go"
	"log"
	"net/http"
	"time"

	httptransport "github.com/go-kit/kit/transport/http"
)

var logger = apilog.New()

func main() {
	// Repository
	userRepository := userrepository.NewUserRepository()

	// Service
	userService := userservices.NewUserService(userRepository)

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(userhttp.EncodeError),
	}

	getUserByIdEndpoint := httptransport.NewServer(
		userhttp.MakeGetByIdEndpoint(userService),
		userhttp.DecodeIdFromUrl,
		userhttp.EncodeResponse,
		serverOptions...,
	)

	router := mux.NewRouter()
	router.Use(apiscope.LifecycleCtxSetup())

	RegisterEndpoint(router, getUserByIdEndpoint, "/users/{id}", http.MethodGet)
	go RegisterConsumer("kafka-python-topic", 0, "tcp", "127.0.0.1:9092")

	if err := http.ListenAndServe(":8080", router); err != nil {
		logger.Error(fmt.Sprintf("Shutdown %s", err.Error()))
	}
}

func RegisterEndpoint(router *mux.Router, server *httptransport.Server, path string, method string) {
	logger.Info(fmt.Sprintf("Registered -> Method:%s Path:%s", method, path))
	router.Handle(path, server).Methods(method)
}

func RegisterConsumer(topic string, partition int, protocol string, address string) {
	logger.Info(fmt.Sprintf("Subscribed Protocol: %s, Address:%s, topic: %s, partition: %d", protocol, address, topic, partition))

	conn, err := kafka.DialLeader(context.Background(), protocol, address, topic, partition)
	if err != nil {
		log.Fatal("failed to dial leader:", err)
	}

	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	batch := conn.ReadBatch(10e3, 1e6) // fetch 10KB min, 1MB max

	b := make([]byte, 10e3) // 10KB max per message
	for {
		_, _ = batch.Read(b)
		fmt.Println(string(b))
	}

	/*if err := batch.Close(); err != nil {
		log.Fatal("failed to close batch:", err)
	}

	if err := conn.Close(); err != nil {
		log.Fatal("failed to close connection:", err)
	}*/
}
