package logger

import (
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

func newZapCore(conf Config) zapcore.Core {
	stdout := zapcore.AddSync(os.Stdout)
	consoleCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(newJsonEncoderConfig()),
		stdout,
		zap.DebugLevel,
	)

	fileCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(newJsonEncoderConfig()),
		zapcore.AddSync(&lumberjack.Logger{
			Filename:   conf.FilePath,
			MaxSize:    10,
			MaxBackups: 3,
			MaxAge:     7,
		}),
		zap.ErrorLevel,
	)

	return zapcore.NewTee(consoleCore, fileCore)
}

func newJsonEncoderConfig() zapcore.EncoderConfig {
	return zapcore.EncoderConfig{
		TimeKey:        "timestamp",
		LevelKey:       "level",
		NameKey:        "name",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "msg",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalLevelEncoder,
		EncodeTime:     zapcore.ISO8601TimeEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
	}
}
