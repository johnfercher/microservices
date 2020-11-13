package apilog

import (
	"context"
	"github.com/johnfercher/microservices/userapi/pkg/api"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Info(ctx context.Context, message string, fields ...zap.Field) {
	logger := api.GetContextLogger(ctx)

	logger.Info(message, fields...)
}

func Error(ctx context.Context, message string, fields ...zap.Field) {
	logger := api.GetContextLogger(ctx)

	logger.Error(message, fields...)
}

func Warn(ctx context.Context, message string, fields ...zap.Field) {
	logger := api.GetContextLogger(ctx)

	logger.Warn(message, fields...)
}

func New() *zap.Logger {
	/*config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	zapLogger, err := config.Build()*/

	lvl := zap.NewAtomicLevelAt(zap.InfoLevel)

	encodeConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,
		EncodeTime:     zapcore.RFC3339NanoTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	var options []zap.Option

	options = append(options, zap.AddCaller(), zap.AddStacktrace(zap.ErrorLevel), zap.AddCallerSkip(2))

	encoder := zapcore.NewJSONEncoder(encodeConfig)
	return zap.New(zapcore.NewCore(encoder, NewStdoutWriter(), lvl), options...)
}
