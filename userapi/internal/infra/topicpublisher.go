package infra

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/johnfercher/microservices/userapi/pkg/api/apierror"
	"github.com/segmentio/kafka-go"
	"go.uber.org/zap"
	"net/http"
)

const (
	idGenerationError         string = "id_generation_error"
	cannotMarshallObjectError string = "cannot_marshall_object_error"
	cannotPublishMessageError string = "cannot_publish_message_error"
	eventEmptyError           string = "event_empty_error"
)

type TopicPublisher interface {
	Publish(ctx context.Context, event string, message interface{}) apierror.ApiError
}

type topicPublisher struct {
	writer *kafka.Writer
}

func NewTopicPublisher(kafkaUrl string, topic string) *topicPublisher {
	return &topicPublisher{kafka.NewWriter(kafka.WriterConfig{
		Brokers:  []string{kafkaUrl},
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	})}
}

func (self *topicPublisher) Publish(ctx context.Context, event string, message interface{}) apierror.ApiError {
	if event == "" {
		apiErr := apierror.New(ctx, eventEmptyError, http.StatusInternalServerError)

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	id, err := uuid.NewRandom()
	if err != nil {
		apiErr := apierror.New(ctx, idGenerationError, http.StatusInternalServerError).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	bytesMessage, err := json.Marshal(message)
	if err != nil {
		apiErr := apierror.New(ctx, cannotMarshallObjectError, http.StatusInternalServerError).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	msg := kafka.Message{
		Key:   []byte(id.String()),
		Value: bytesMessage,
		Headers: []kafka.Header{{
			Key: fmt.Sprintf("event:%s", event),
		}},
	}

	err = self.writer.WriteMessages(ctx, msg)
	if err != nil {
		apiErr := apierror.New(ctx, cannotPublishMessageError, http.StatusInternalServerError).
			AppendFields(zap.String("err", err.Error()))

		apierror.Log(ctx, apiErr)
		return apiErr
	}

	return nil
}
