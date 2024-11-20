package main

import (
	"context"
	"log"
	"time"

	"github.com/DucTran999/shared-pkg/logger"
)

func main() {
	conf := logger.Config{
		Environment: logger.Production,
		LogToFile:   true,
		FilePath:    "logs/app.log",
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		log.Fatalln("Init logger ERR", err)
	}

	i := 0
	for i = 0; i < 10000; i++ {
		go func(logger.ILogger) {
			logInst.FromContext(context.Background()).Error("example error log")
		}(logInst)
	}

	defer logInst.Sync()
	time.Sleep(time.Second * 10)
}
