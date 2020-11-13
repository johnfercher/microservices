package api

import (
	"context"
	"go.uber.org/zap"
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
func AddContextLogger(ctx context.Context, logger *zap.Logger) context.Context {
	return AddContextField(ctx, CtxLogger, logger)
}

func GetContextLogger(ctx context.Context) *zap.Logger {
	logger, ok := GetContextField(ctx, CtxLogger).(*zap.Logger)
	if !ok {
		return nil
	}

	return logger
}
