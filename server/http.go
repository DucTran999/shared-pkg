package server

import (
	"context"
	"fmt"

	"net/http"
	"os/signal"
	"syscall"
	"time"

	"github.com/rs/zerolog/log"
)

type FreeResourceFunc func(ctx context.Context) error

type HttpServer interface {
	Start() error
	Stop(ctx context.Context) error
}

type httpServer struct {
	server *http.Server
}

func (s *httpServer) Start() error {
	log.Info().Msgf("start server on %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *httpServer) Stop(ctx context.Context) error {
	log.Info().Msg("http server is shutting down...")

	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server got err: %v", err)
	}

	log.Info().Msg("http server shutdown successfully!")
	return nil
}

// GracefulShutdown handles OS signals and performs a graceful shutdown of the server.
func GracefulShutdown(shutdownTasks ...FreeResourceFunc) {
	const shutdownTimeout = 5 * time.Second

	// Listen for SIGINT or SIGTERM
	shutdownCtx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	<-shutdownCtx.Done()
	log.Info().Msg("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	cleanExit := true
	for _, task := range shutdownTasks {
		if err := task(ctx); err != nil {
			log.Warn().Err(err).Msg("shutdown task error")
			cleanExit = false
		}
	}

	if cleanExit {
		log.Info().Msg("server shut down cleanly")
	} else {
		log.Warn().Msg("server encountered errors during shutdown")
	}
}
