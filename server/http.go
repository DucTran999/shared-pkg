package server

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

type IHttpServer interface {
	Start() error
	Stop() error
}

type httpServer struct {
	server *http.Server
}

func (s *httpServer) Start() error {
	log.Printf("server on %s", s.server.Addr)
	if err := s.server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("start server got err: %v", err)
	}

	return nil
}

func (s *httpServer) Stop() error {
	log.Println("http server is shutting down...")
	if err := s.server.Shutdown(context.Background()); err != nil {
		return fmt.Errorf("shutdown server got err: %v", err)
	}

	log.Println("http server shutdown successfully!")
	return nil
}

// gracefulShutdown handles OS signals and performs a graceful shutdown of the server.
func GracefulShutdown(shutdownTasks ...func() error) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	// Block until we receive a termination signal
	<-quit
	log.Println("Shutting down server...")

	// Create a context with a timeout to enforce graceful shutdown timing
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Track if all tasks exit cleanly
	cleanExit := true

	for _, task := range shutdownTasks {
		if err := task(); err != nil {
			log.Println(err)
			cleanExit = false
		}
	}

	// Wait for the context to finish or timeout
	<-ctx.Done()

	if cleanExit {
		log.Println("Server shut down cleanly")
	} else {
		log.Println("Server encountered errors during shutdown")
	}
}
