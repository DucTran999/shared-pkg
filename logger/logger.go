package logger

import (
	"context"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type logger struct {
	zapLogger   *zap.Logger
	sugarLogger *zap.SugaredLogger
}

func NewLogger(conf Config) (ILogger, error) {
	// Create the base zap core
	core := newZapCore(conf)

	// Apply sampling only for production environment
	if conf.Environment == Production {
		core = zapcore.NewSamplerWithOptions(
			core,
			time.Second, // interval
			100,         // log first 100 entries
			100,         // thereafter log zero entires within the interval
		)
	}

	zapLog := zap.New(
		core,
		zap.AddCaller(),
		zap.AddCallerSkip(1),
		zap.AddStacktrace(zap.ErrorLevel),
	)

	return &logger{
		zapLogger:   zapLog,
		sugarLogger: zapLog.Sugar(),
	}, nil
}

func (l *logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

func (l *logger) Error(msg string, fields ...zap.Field) {
	l.zapLogger.Error(msg, fields...)
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.zapLogger.Panic(msg, fields...)
}

func (l *logger) DPanic(msg string, fields ...zap.Field) {
	l.zapLogger.DPanic(msg, fields...)
}

func (l *logger) Infof(template string, args ...any) {
	l.sugarLogger.Infof(template, args...)
}

func (l *logger) Warnf(template string, args ...any) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *logger) Errorf(template string, args ...any) {
	l.sugarLogger.Errorf(template, args...)
}

func (l *logger) Debugf(template string, args ...any) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *logger) Panicf(template string, args ...any) {
	l.sugarLogger.Panicf(template, args...)
}

func (l *logger) DPanicf(template string, args ...any) {
	l.sugarLogger.DPanicf(template, args...)
}

// FromContext retrieves data from the context and returns a logger with those fields
func (l *logger) FromContext(ctx context.Context) ILogger {
	// Extract the request ID from the context
	requestID, ok := ctx.Value(requestIDKeyCtx).(string)

	// If a non-empty request ID exists, attach it to the logger
	if ok && requestID != "" {
		newLogger := l.zapLogger.With(zap.String(requestIDKeyCtx, requestID))
		return &logger{zapLogger: newLogger, sugarLogger: newLogger.Sugar()}
	}

	// Return the original logger if no valid request ID is found
	return l
}

func (l *logger) Sync() error {
	return l.zapLogger.Sync()
}
