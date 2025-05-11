package logger_test

import (
	"context"
	"errors"
	"log"
	"sync"
	"testing"

	"github.com/DucTran999/shared-pkg/logger"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func Test_LogTypeWithoutRecover(t *testing.T) {
	conf := logger.Config{
		Environment: logger.Production,
		LogToFile:   true,
		FilePath:    "logs/app.log",
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		log.Fatalln("Init logger ERR", err)
	}
	defer func() { _ = logInst.Sync() }()

	var wg sync.WaitGroup
	for range 100 {
		wg.Add(1)
		go func(logger.ILogger) {
			ctx := context.Background()
			newCtx := context.WithValue(ctx, logger.RequestIDKeyCtx, "123456")

			defer wg.Done()
			logInst.FromContext(newCtx).Error("example error log")
			logInst.FromContext(newCtx).Info("example info log")
			logInst.FromContext(newCtx).Debug("example error log")
			logInst.FromContext(newCtx).Warn("example warn log")

			logInst.FromContext(context.Background()).Infof("example info log %s", "test")
			logInst.FromContext(context.Background()).Debugf("example debug log %s", "test")
			logInst.FromContext(context.Background()).Errorf("example error log %s", "test")
			logInst.FromContext(context.Background()).Warnf("example warn log %s", "test")
		}(logInst)
	}
	wg.Wait()
}

func Test_Panic(t *testing.T) {
	conf := logger.Config{
		Environment: logger.Production,
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		t.Fatalf("Init logger ERR=%v", err)
	}
	defer func() { _ = logInst.Sync() }()

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example panic log", zap.Any("stack", r))
			panicOccurred = true
		}
		require.True(t, panicOccurred, "Expected panic did not occur in Staging environment")
	}()

	logInst.Panic("example panic log", zap.String("stack", "stack trace"))
}

func Test_Panicf(t *testing.T) {
	conf := logger.Config{
		Environment: logger.Staging,
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		t.Fatalf("Init logger ERR=%v", err)
	}
	defer func() { _ = logInst.Sync() }()

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example panic log", zap.Any("stack", r))
			panicOccurred = true
		}
		require.True(t, panicOccurred, "Expected panic did not occur in Staging environment")
	}()

	logInst.Panicf("example panic log %v", errors.New("panic test"))
}

func Test_DPanic(t *testing.T) {
	conf := logger.Config{
		Environment: logger.Staging,
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		t.Fatalf("Init logger ERR=%v", err)
	}
	defer func() { _ = logInst.Sync() }()

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example dpanic log", zap.Any("stack", r))
			panicOccurred = true
		}
		require.True(t, panicOccurred, "Expected panic did not occur in Staging environment")
	}()

	logInst.DPanic("example dpanic log")
}

func Test_DPanicf(t *testing.T) {
	conf := logger.Config{
		Environment: logger.Staging,
	}

	logInst, err := logger.NewLogger(conf)
	if err != nil {
		t.Fatalf("Init logger ERR=%v", err)
	}
	defer func() { _ = logInst.Sync() }()

	panicOccurred := false
	defer func() {
		if r := recover(); r != nil {
			logInst.Error("example dpanicd log", zap.Any("stack", r))
			panicOccurred = true
		}
		require.True(t, panicOccurred, "Expected panic did not occur in Staging environment")
	}()

	logInst.DPanicf("example dpanicf log err:%v", errors.New("panic test"))
}
