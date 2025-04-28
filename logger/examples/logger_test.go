package main

import (
	"context"
	"log"
	"testing"

	"github.com/DucTran999/shared-pkg/logger"
)

func Test_Logging(t *testing.T) {
	conf := logger.Config{
		Environment: logger.Production,
		LogToFile:   true,
		FilePath:    "logs/app.log",
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		log.Fatalln("Init logger ERR", err)
	}
	defer logInst.Sync()

	for range 10000 {
		go func(logger.ILogger) {
			logInst.FromContext(context.Background()).Error("example error log")
		}(logInst)
	}
}
