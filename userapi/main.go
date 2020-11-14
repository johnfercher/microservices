package main

import (
	"context"
	"fmt"
	"log"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func getKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
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
k
func main() {
	// get kafka reader using environment variables.
	kafkaURL := "localhost:9092"
	topic := "kafka-topic-test"
	groupID := "group1"

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
}

/*var logger = apilog.New()

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
	go RegisterConsumer("my-topic", 0, "tcp", "127.0.0.1:9092")

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

	kafka.Consu

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
	}
}*/