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
	Stop(ctx context.Context) error
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

func (s *httpServer) Stop(ctx context.Context) error {
	s.logger.Info().Msg("http server is shutting down...")

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server got err: %v", err)
	}

	s.logger.Info().Msg("http server shutdown successfully!")
	return nil
}

// GracefulShutdown handles OS signals and performs a graceful shutdown of the server.
func GracefulShutdown(shutdownTasks ...func(ctx context.Context) error) {
	const shutdownTimeout = 5 * time.Second
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()

	// Listen for SIGINT or SIGTERM
	shutdownCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdownCtx.Done()
	logger.Info().Msg("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	cleanExit := true
	for _, task := range shutdownTasks {
		if err := task(ctx); err != nil {
			logger.Warn().Err(err).Msg("shutdown task error")
			cleanExit = false
		}
	}

	if cleanExit {
		logger.Info().Msg("Server shut down cleanly")
	} else {
		logger.Warn().Msg("Server encountered errors during shutdown")
	}
}
