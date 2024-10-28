package server

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
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

	return &ginServer{
		httpServer: &httpServer{server},
	}
}
