package logger

import (
	"context"
	"fmt"
	"os"
	"path"
	"runtime"

	"github.com/sirupsen/logrus"
)

var Logger *logrus.Logger

// Context key for request ID
type contextKey string

const RequestIDKey contextKey = "request_id"

func init() {
	Logger = logrus.New()
	Logger.SetOutput(os.Stdout)

	Logger.SetFormatter(&logrus.JSONFormatter{
		CallerPrettyfier: func(f *runtime.Frame) (string, string) {
			filename := path.Base(f.File)
			return "", fmt.Sprintf("%s:%d", filename, f.Line)
		},
	})
	Logger.SetReportCaller(true)

	level := os.Getenv("LOG_LEVEL")
	switch level {
	case "DEBUG":
		Logger.SetLevel(logrus.DebugLevel)
	case "WARN":
		Logger.SetLevel(logrus.WarnLevel)
	case "ERROR":
		Logger.SetLevel(logrus.ErrorLevel)
	case "FATAL":
		Logger.SetLevel(logrus.FatalLevel)
	default:
		Logger.SetLevel(logrus.InfoLevel)
	}
}

func GetLogger() *logrus.Logger {
	return Logger
}

// WithRequestID returns a logger entry with request ID from context
func WithRequestID(ctx context.Context) *logrus.Entry {
	if requestID := ctx.Value(RequestIDKey); requestID != nil {
		return Logger.WithField("request_id", requestID)
	}
	return Logger.WithField("request_id", "")
}

func WithField(key string, value interface{}) *logrus.Entry {
	return Logger.WithField(key, value)
}

func WithFields(fields map[string]interface{}) *logrus.Entry {
	return Logger.WithFields(logrus.Fields(fields))
}

func WithError(err error) *logrus.Entry {
	return Logger.WithError(err)
}

// Context-aware logging functions
func InfoCtx(ctx context.Context, msg string) {
	WithRequestID(ctx).Info(msg)
}

func DebugCtx(ctx context.Context, msg string) {
	WithRequestID(ctx).Debug(msg)
}

func WarnCtx(ctx context.Context, msg string) {
	WithRequestID(ctx).Warn(msg)
}

func ErrorCtx(ctx context.Context, msg string) {
	WithRequestID(ctx).Error(msg)
}

func FatalCtx(ctx context.Context, msg string) {
	WithRequestID(ctx).Fatal(msg)
}

// Legacy functions (without context)
func Info(msg string) {
	Logger.Info(msg)
}

func Debug(msg string) {
	Logger.Debug(msg)
}

func Warn(msg string) {
	Logger.Warn(msg)
}

func Error(msg string) {
	Logger.Error(msg)
}

func Fatal(msg string) {
	Logger.Fatal(msg)
}
