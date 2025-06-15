package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

var Logger *zap.Logger

func init() {
	config := zap.NewProductionConfig()
	config.Encoding = "json"
	config.EncoderConfig.TimeKey = "timestamp"
	config.EncoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	config.EncoderConfig.CallerKey = "caller"
	config.EncoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	config.OutputPaths = []string{"stdout"}
	config.ErrorOutputPaths = []string{"stdout"}

	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "DEBUG":
		config.Level = zap.NewAtomicLevelAt(zapcore.DebugLevel)
	case "WARN":
		config.Level = zap.NewAtomicLevelAt(zapcore.WarnLevel)
	case "ERROR":
		config.Level = zap.NewAtomicLevelAt(zapcore.ErrorLevel)
	case "FATAL":
		config.Level = zap.NewAtomicLevelAt(zapcore.FatalLevel)
	default:
		config.Level = zap.NewAtomicLevelAt(zapcore.InfoLevel)
	}

	var err error
	Logger, err = config.Build()
	if err != nil {
		panic(err)
	}
}

func GetLogger() *zap.Logger {
	return Logger
}

func WithField(key string, value interface{}) *zap.Logger {
	return Logger.With(zap.Any(key, value))
}

func WithFields(fields map[string]interface{}) *zap.Logger {
	zapFields := make([]zap.Field, 0, len(fields))
	for k, v := range fields {
		zapFields = append(zapFields, zap.Any(k, v))
	}
	return Logger.With(zapFields...)
}

func WithError(err error) *zap.Logger {
	return Logger.With(zap.Error(err))
}

func Info(msg string, fields ...zap.Field) {
	Logger.Info(msg, fields...)
}

func Debug(msg string, fields ...zap.Field) {
	Logger.Debug(msg, fields...)
}

func Warn(msg string, fields ...zap.Field) {
	Logger.Warn(msg, fields...)
}

func Error(msg string, fields ...zap.Field) {
	Logger.Error(msg, fields...)
}

func Fatal(msg string, fields ...zap.Field) {
	Logger.Fatal(msg, fields...)
}
