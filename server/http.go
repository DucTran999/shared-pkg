package server

import (
	"context"
	"fmt"
	"net"

	"net/http"
	"time"

	"github.com/rs/zerolog/log"
)

type HttpServer interface {
	Start() error
	Stop(ctx context.Context) error
}

type httpServer struct {
	server *http.Server
}

func (s *httpServer) Start() error {
	go func() {
		for {
			// Wait for the server to be ready
			time.Sleep(500 * time.Millisecond)

			// Attempt to ping the server to check if it's ready
			err := s.tcpPing(s.server.Addr, 5*time.Second)
			if err != nil {
				log.Warn().Err(err).Msgf("tcp ping to %s failed, retrying...", s.server.Addr)
				continue
			}

			// If ping is successful, log the server address and break the loop
			log.Info().Msgf("start server on %s", s.server.Addr)
			break
		}
	}()

	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return err
	}

	return nil
}

func (s *httpServer) Stop(ctx context.Context) error {
	if err := s.server.Shutdown(ctx); err != nil {
		return fmt.Errorf("shutdown server got err: %v", err)
	}

	log.Info().Msg("http server shutdown successfully")
	return nil
}

func (s *httpServer) tcpPing(address string, timeout time.Duration) error {
	conn, err := net.DialTimeout("tcp", address, timeout)
	if err != nil {
		return err
	}

	defer conn.Close()
	return nil
}
