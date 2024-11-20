package logger

import (
	"context"

	"go.uber.org/zap"
)

type Config struct {
	Environment string
	LogToFile   bool
	FilePath    string
}

type ILogger interface {
	// Standard logging methods
	Debug(msg string, fields ...zap.Field)
	Info(msg string, fields ...zap.Field)
	Warn(msg string, fields ...zap.Field)
	Error(msg string, fields ...zap.Field)
	Fatal(msg string, fields ...zap.Field)
	Panic(msg string, fields ...zap.Field)
	DPanic(msg string, fields ...zap.Field)

	// Formatted logging methods
	Debugf(template string, args ...any)
	Infof(template string, args ...any)
	Warnf(template string, args ...any)
	Errorf(template string, args ...any)
	Fatalf(template string, args ...any)
	Panicf(template string, args ...any)
	DPanicf(template string, args ...any)

	FromContext(context.Context) ILogger
	Sync() error
}
