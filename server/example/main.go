package main

import (
	"github.com/DucTran999/shared-pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	httpServer := server.NewGinHttpServer(router, "localhost", 8000)

	go func() {
		if err := httpServer.Start(); err != nil {
			log.Fatal().Msgf("start http server got ERR=%v", err)
		}
	}()

	server.GracefulShutdown(httpServer.Stop)
}
