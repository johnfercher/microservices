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

/*func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		StartOffset: kafka.LastOffset,
	})
}

func main() {
	// get kafka reader using environment variables.
	kafkaURL := "localhost:9092"
	topic := "kafka-topic-test"
	groupID := "group1"
k
	reader := getKafkaReader(kafkaURL, topic, groupID)

	defer reader.Close()

	fmt.Println("start consuming ... !!")
	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		fmt.Printf("message at topic:%v partition:%v offset:%v	%s = %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}*/

var logger = apilog.New()

func main() {
	// MySql
	mysql, err := infra.GetMysqlConnection()
	if err != nil {
		panic(err)
	}

	// Repository
	userRepository := userrepository.NewUserRepository(mysql)

	// Service
	userService := userservice.NewUserService(userRepository)

	serverOptions := []httptransport.ServerOption{
		httptransport.ServerErrorEncoder(userhttp.EncodeError),
	}

	router := mux.NewRouter()
	router.Use(apiscope.LifecycleCtxSetup())

	RegisterEndpoint(router, "/users/{id}", http.MethodGet, httptransport.NewServer(
		userhttp.MakeGetByIdEndpoint(userService),
		userhttp.DecodeIdFromUrl,
		userhttp.EncodeResponse,
		serverOptions...,
	))

	RegisterEndpoint(router, "/users", http.MethodPost, httptransport.NewServer(
		userhttp.MakeCreateEndpoint(userService),
		userhttp.DecodeCreateUserRequestFromBody,
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
