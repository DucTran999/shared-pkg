package server

import (
	"context"
	"fmt"

	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog"
)

type HttpServer interface {
	Start() error
	Stop() error
}

type httpServer struct {
	server *http.Server
	logger *zerolog.Logger
}

func (s *httpServer) Start() error {
	s.logger.Info().Msgf("start server on %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *httpServer) Stop() error {
	s.logger.Info().Msg("http server is shutting down...")
	if err := s.server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("shutdown server got err: %v", err)
	}

	s.logger.Info().Msg("http server shutdown successfully!")
	return nil
}

// gracefulShutdown handles OS signals and performs a graceful shutdown of the server.
func GracefulShutdown(shutdownTasks ...func() error) {
	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
	logger := zerolog.New(output).With().Timestamp().Logger()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a termination signal
	<-quit
	logger.Info().Msg("Shutting down server...")

	// Create a context with a timeout to enforce graceful shutdown timing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Track if all tasks exit cleanly
	cleanExit := true

	for _, task := range shutdownTasks {
		if err := task(); err != nil {
			logger.Warn().Msg(err.Error())
			cleanExit = false
		}
	}

	// Wait for the context to finish or timeout
	<-ctx.Done()

	if cleanExit {
		logger.Info().Msg("Server shut down cleanly")
	} else {
		logger.Info().Msg("Server encountered errors during shutdown")
	}
}
