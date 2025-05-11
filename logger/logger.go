package logger

import (
	"context"
	"fmt"
	"runtime"
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
		// zap.AddCallerSkip(1),
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
	l.logWithStack(zapcore.ErrorLevel, msg, fields...)
}

func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.PanicLevel, msg, fields...)
}

func (l *logger) DPanic(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.DPanicLevel, msg, fields...)
}

func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.FatalLevel, msg, fields...)
}

func (l *logger) Infof(template string, args ...any) {
	l.sugarLogger.Infof(template, args...)
}

func (l *logger) Warnf(template string, args ...any) {
	l.sugarLogger.Warnf(template, args...)
}

func (l *logger) Errorf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.ErrorLevel, msg)
}

func (l *logger) Debugf(template string, args ...any) {
	l.sugarLogger.Debugf(template, args...)
}

func (l *logger) Panicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.PanicLevel, msg)
}

func (l *logger) DPanicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.DPanicLevel, msg)
}

func (l *logger) Fatalf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.FatalLevel, msg)
}

// FromContext retrieves data from the context and returns a logger with those fields
func (l *logger) FromContext(ctx context.Context) ILogger {
	// Extract the request ID from the context
	requestID, ok := ctx.Value(RequestIDKeyCtx).(string)

	// If a non-empty request ID exists, attach it to the logger
	if ok && requestID != "" {
		newLogger := l.zapLogger.With(zap.String(RequestIDKeyCtx, requestID))
		return &logger{zapLogger: newLogger, sugarLogger: newLogger.Sugar()}
	}

	// Return the original logger if no valid request ID is found
	return l
}

func (l *logger) Sync() error {
	return l.zapLogger.Sync()
}

// logWithStack logs a message with optional stack trace information for error-level logs and above.
func (l *logger) logWithStack(level zapcore.Level, msg string, fields ...zap.Field) {
	// For error and above, capture caller information
	pc, file, line, ok := runtime.Caller(2)
	if ok {
		if fn := runtime.FuncForPC(pc); fn != nil {
			stacktrace := zap.Dict("source",
				zap.String("path", file),
				zap.Int("line", line),
				zap.String("func", fn.Name()),
			)
			fields = append(fields, stacktrace)
		}
	}

	// Log with stack trace
	if ce := l.zapLogger.Check(level, msg); ce != nil {
		ce.Write(fields...)
	}
}
