package apilog

import (
	"context"
	"github.com/johnfercher/microservices/userapi/pkg/api/apifields"
	"github.com/johnfercher/microservices/userapi/pkg/api/apiglobal"
	"github.com/sirupsen/logrus"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
	"os"
)

func Info(ctx context.Context, message string, fields ...apifields.Field) {
	logger := GetContextLogger(ctx)

	logger.Info(message, fields...)
}

func Error(ctx context.Context, message string, fields ...apifields.Field) {
	logger := GetContextLogger(ctx)

	logger.Error(message, fields...)
}

func Warn(ctx context.Context, message string, fields ...apifields.Field) {
	logger := GetContextLogger(ctx)

	logger.Warn(message, fields...)
}

type ApiLogger interface {
	Error(message string, fields ...apifields.Field)
	Info(message string, fields ...apifields.Field)
	Warn(message string, fields ...apifields.Field)
	WithKeyValue(key string, value interface{}) ApiLogger
}

type apiLogger struct {
	logger   *logrus.Logger
	keyValue map[string]interface{}
}

func New() ApiLogger {
	var logger = &logrus.Logger{
		Out:   os.Stderr,
		Hooks: make(logrus.LevelHooks),
		Level: logrus.DebugLevel,
		Formatter: &logrus.JSONFormatter{
			FieldMap: logrus.FieldMap{
				logrus.FieldKeyTime:  "@timestamp",
				logrus.FieldKeyLevel: "log.level",
				logrus.FieldKeyMsg:   "message",
				logrus.FieldKeyFunc:  "function.name", // non-ECS
			},
		},
	}

	hook := graylog.NewGraylogHook(apiglobal.GetGlobalConfig().Logstash.Url, map[string]interface{}{})
	logger.AddHook(hook)

	return &apiLogger{
		logger:   logger,
		keyValue: make(map[string]interface{}),
	}
}

func (self *apiLogger) Error(message string, fields ...apifields.Field) {
	logKeyValues := make(map[string]interface{})

	for key, value := range self.keyValue {
		logKeyValues[key] = value
	}

	for _, field := range fields {
		logKeyValues[field.Key] = field.Value
	}

	self.logger.WithFields(logKeyValues).Error(message)
}

func (self *apiLogger) Info(message string, fields ...apifields.Field) {
	logKeyValues := make(map[string]interface{})

	for key, value := range self.keyValue {
		logKeyValues[key] = value
	}

	for _, field := range fields {
		logKeyValues[field.Key] = field.Value
	}

	self.logger.WithFields(logKeyValues).Info(message)
}

func (self *apiLogger) Warn(message string, fields ...apifields.Field) {
	logKeyValues := make(map[string]interface{})

	for key, value := range self.keyValue {
		logKeyValues[key] = value
	}

	for _, field := range fields {
		logKeyValues[field.Key] = field.Value
	}

	self.logger.WithFields(logKeyValues).Warn(message)
}

func (self *apiLogger) WithKeyValue(key string, value interface{}) ApiLogger {
	self.keyValue[key] = value
	return self
}
