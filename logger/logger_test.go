package logger_test

import (
	"context"
	"log"
	"sync"
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

	var wg sync.WaitGroup
	for range 10000 {
		wg.Add(1)
		go func(logger.ILogger) {
			defer wg.Done()
			logInst.FromContext(context.Background()).Error("example error log")
		}(logInst)
	}
	wg.Wait()
}
