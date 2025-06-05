package logger

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// logger is an implementation of the ILogger interface,
// wrapping both a structured zap.Logger and a sugared logger for formatted output.
type logger struct {
	zapLogger   *zap.Logger
	sugarLogger *zap.SugaredLogger
}

// NewLogger creates and returns a new ILogger instance configured
// according to the provided Config.
//
// The logger uses a custom zapcore and applies sampling in production
// environments to limit the volume of logs. In development mode,
// it enables development-specific options (like DPanic triggering).
//
// Parameters:
//   - conf: Config struct defining the logging environment, output settings, etc.
//
// Returns:
//   - ILogger: a structured and sugared logger instance
//   - error: if any error occurs during logger initialization
func NewLogger(conf Config) (ILogger, error) {
	// Create the base zap core
	core := newZapCore(conf)

	// Apply sampling only for production environment
	if conf.Environment == Production {
		core = zapcore.NewSamplerWithOptions(
			core,
			time.Second, // interval
			100,         // log first 100 entries
			100,         // thereafter log zero entries within the interval
		)
	}

	zapLog := zap.New(core)
	if conf.Environment == Development {
		zapLog = zapLog.WithOptions(zap.Development())
	}

	return &logger{
		zapLogger:   zapLog,
		sugarLogger: zapLog.Sugar(),
	}, nil
}

// Info logs a message at InfoLevel.
func (l *logger) Info(msg string, fields ...zap.Field) {
	l.zapLogger.Info(msg, fields...)
}

// Warn logs a message at WarnLevel.
func (l *logger) Warn(msg string, fields ...zap.Field) {
	l.zapLogger.Warn(msg, fields...)
}

// Error logs a message at ErrorLevel and includes a stack trace.
func (l *logger) Error(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.ErrorLevel, msg, fields...)
}

// Debug logs a message at DebugLevel.
func (l *logger) Debug(msg string, fields ...zap.Field) {
	l.zapLogger.Debug(msg, fields...)
}

// Panic logs a message at PanicLevel and then panics. Includes a stack trace.
func (l *logger) Panic(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.PanicLevel, msg, fields...)
}

// DPanic logs a message at DPanicLevel.
// In development, it panics. In production, it only logs. Includes a stack trace.
func (l *logger) DPanic(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.DPanicLevel, msg, fields...)
}

// Fatal logs a message at FatalLevel and then calls os.Exit(1). Includes a stack trace.
func (l *logger) Fatal(msg string, fields ...zap.Field) {
	l.logWithStack(zapcore.FatalLevel, msg, fields...)
}

// Infof logs a formatted message at InfoLevel.
func (l *logger) Infof(template string, args ...any) {
	l.sugarLogger.Infof(template, args...)
}

// Warnf logs a formatted message at WarnLevel.
func (l *logger) Warnf(template string, args ...any) {
	l.sugarLogger.Warnf(template, args...)
}

// Errorf logs a formatted message at ErrorLevel and includes a stack trace.
func (l *logger) Errorf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.ErrorLevel, msg)
}

// Debugf logs a formatted message at DebugLevel.
func (l *logger) Debugf(template string, args ...any) {
	l.sugarLogger.Debugf(template, args...)
}

// Panicf logs a formatted message at PanicLevel and panics. Includes a stack trace.
func (l *logger) Panicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.PanicLevel, msg)
}

// DPanicf logs a formatted message at DPanicLevel. Includes a stack trace.
// Panics in development mode.
func (l *logger) DPanicf(template string, args ...any) {
	msg := fmt.Sprintf(template, args...)
	l.logWithStack(zapcore.DPanicLevel, msg)
}

// Fatalf logs a formatted message at FatalLevel and exits the program. Includes a stack trace.
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
		newLogger := l.zapLogger.With(zap.String(string(RequestIDKeyCtx), requestID))
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
