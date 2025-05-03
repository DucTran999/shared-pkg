package server

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type ginServer struct {
	*httpServer
}

func NewGinHttpServer(router *gin.Engine, host string, port int) *ginServer {
	addr := fmt.Sprintf("%s:%d", host, port)
	server := &http.Server{
		Addr:    addr,
		Handler: router,
	}

	output := zerolog.ConsoleWriter{Out: os.Stdout, TimeFormat: time.RFC3339, NoColor: false}
	logger := zerolog.New(output).With().Timestamp().Logger()

	return &ginServer{
		httpServer: &httpServer{
			server: server,
			logger: &logger,
		},
	}
}
