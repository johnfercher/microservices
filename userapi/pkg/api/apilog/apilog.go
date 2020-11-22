package apilog

import (
	"context"
	"github.com/johnfercher/microservices/userapi/pkg/api"
	"github.com/sirupsen/logrus"
	"go.uber.org/zap"
	graylog "gopkg.in/gemnasium/logrus-graylog-hook.v2"
	"os"
)

func Info(ctx context.Context, message string, fields ...zap.Field) {
	logger := api.GetContextLogger(ctx)

	logger.Info(message /*, fields...*/)
}

func Error(ctx context.Context, message string, fields ...zap.Field) {
	logger := api.GetContextLogger(ctx)

	logger.Error(message /*, fields...*/)
}

func Warn(ctx context.Context, message string, fields ...zap.Field) {
	logger := api.GetContextLogger(ctx)

	logger.Warn(message /*, fields...*/)
}

func New(logstashServer string) *logrus.Logger {
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

	hook := graylog.NewGraylogHook(logstashServer, map[string]interface{}{})
	logger.AddHook(hook)

	return logger

	/*config := zap.NewProductionConfig()
	config.OutputPaths = []string{"stdout"}
	zapLogger, err := config.Build()

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
	stdErr := zapcore.Lock(zapcore.AddSync(os.Stderr))

	encoder := zapcore.NewJSONEncoder(encodeConfig)
	logger := zap.New(zapcore.NewCore(encoder, stdErr, lvl), options...)

	return logger*/
}
