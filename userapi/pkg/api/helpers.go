package api

import (
	"context"
	"github.com/sirupsen/logrus"
)

// Generic
func AddContextField(ctx context.Context, key string, value interface{}) context.Context {
	return context.WithValue(ctx, key, value)
}

func GetContextField(ctx context.Context, key string) interface{} {
	return ctx.Value(key)
}

func GetContextStringField(ctx context.Context, key string) string {
	stringField, ok := GetContextField(ctx, key).(string)
	if !ok {
		return ""
	}

	return stringField
}

// Logging
func AddContextLogger(ctx context.Context, logger *logrus.Logger) context.Context {
	return AddContextField(ctx, CtxLogger, logger)
}

func GetContextLogger(ctx context.Context) *logrus.Logger {
	logger, ok := GetContextField(ctx, CtxLogger).(*logrus.Logger)
	if !ok {
		return nil
	}

	return logger
}
