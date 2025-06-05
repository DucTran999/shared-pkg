//go:build ignore
// +build ignore

package main

import (
	"github.com/DucTran999/shared-pkg/logger"
	"go.uber.org/zap"
)

func main() {
	loggerInst, _ := logger.NewLogger(logger.Config{
		Environment: logger.Production,
		LogToFile:   false,
		FilePath:    "logs/app.log",
	})
	defer loggerInst.Sync()

	loggerInst.Error("example error log", zap.Int("user_id", 166))
}
