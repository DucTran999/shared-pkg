package server

import (
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var (
	ErrPortRequired = errors.New("port error: must be positive integer")
)

// ServerConfig holds the configuration settings for the HTTP server.
type ServerConfig struct {
	// Host is the address or hostname of the server.
	// This will be combined with the Port field to form the full address.
	Host string

	// Port specifies the port on which the server will listen.
	// It must be a positive integer (e.g., 8080).
	Port int

	// ReadHeaderTimeout defines the maximum duration the server will wait
	// for the client to send the headers of the request.
	// If this time is exceeded, the server will return a timeout error.
	ReadHeaderTimeout time.Duration

	// ReadTimeout defines the maximum duration the server will wait for the entire
	// request to be read after receiving the request headers.
	// If the request body takes longer than this duration, the request will be rejected.
	ReadTimeout time.Duration

	// WriteTimeout defines the maximum duration the server will wait for a response
	// to be sent to the client after it has been generated.
	// If it takes longer than this, the server will return a timeout error to the client.
	// WriteTimeout = Handle + Write
	WriteTimeout time.Duration

	// IdleTimeout defines the maximum time that the connection will be kept open
	// while waiting for another request after completing the previous request.
	// If the idle time exceeds this value, the connection will be closed.
	IdleTimeout time.Duration
}

type ginServer struct {
	*httpServer
}

// NewGinHttpServer initializes and returns a new Gin HTTP server
// with the provided router and configuration settings.
func NewGinHttpServer(router *gin.Engine, config ServerConfig) (*ginServer, error) {
	// Validate and apply defaults to the provided server configuration
	cleanConfig, err := setupConfig(config)
	if err != nil {
		return nil, err
	}

	addr := fmt.Sprintf("%s:%d", cleanConfig.Host, cleanConfig.Port)
	server := &http.Server{
		Addr:              addr,
		Handler:           router,
		ReadHeaderTimeout: cleanConfig.ReadHeaderTimeout,
		ReadTimeout:       cleanConfig.ReadTimeout,
		WriteTimeout:      cleanConfig.WriteTimeout,
		IdleTimeout:       cleanConfig.IdleTimeout,
	}

	return &ginServer{
		httpServer: &httpServer{
			server: server,
		},
	}, nil
}

// setupConfig applies default values and validates the server configuration.
func setupConfig(config ServerConfig) (*ServerConfig, error) {
	cleanConfig := config

	// Port must be specified and greater than 0
	if cleanConfig.Port <= 0 {
		return nil, ErrPortRequired
	}

	// Set default idle timeout if not provided
	if cleanConfig.IdleTimeout == 0 {
		cleanConfig.IdleTimeout = 60 * time.Second
	}

	// Set default read timeout if not provided
	if cleanConfig.ReadTimeout == 0 {
		cleanConfig.ReadTimeout = 10 * time.Second
	}

	// Set default read header timeout if not provided
	if cleanConfig.ReadHeaderTimeout == 0 {
		cleanConfig.ReadHeaderTimeout = 500 * time.Millisecond
	}

	// Set default write timeout if not provided
	if cleanConfig.WriteTimeout == 0 {
		cleanConfig.WriteTimeout = 10 * time.Second
	}

	return &cleanConfig, nil
}
