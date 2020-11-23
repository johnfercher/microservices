package apilog

import (
	"context"
)

const (
	HeaderRequestId string = "X-Request-Id"

	CtxRequestId     string = "request.id"
	CtxRequestMethod string = "request.method"
	CtxRequestHost   string = "request.host"
	CtxContentLength string = "request.content_length"
	CtxTlsVersion    string = "request.tls_version"
	CtxRequestPath   string = "request.path"
	CtxRequestQuery  string = "request.query"
	CtxLogger        string = "logger"
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

// Logger
func AddContextLogger(ctx context.Context, logger ApiLogger) context.Context {
	return AddContextField(ctx, CtxLogger, logger)
}

func GetContextLogger(ctx context.Context) ApiLogger {
	logger, ok := GetContextField(ctx, CtxLogger).(ApiLogger)
	if !ok {
		return nil
	}

	return logger
}
