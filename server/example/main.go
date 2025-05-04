package main

import (
	"github.com/DucTran999/shared-pkg/server"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
)

func main() {
	gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	httpServer, err := server.NewGinHttpServer(router, server.ServerConfig{
		Host: "localhost",
		Port: 8080,
	})
	if err != nil {
		log.Fatal().Msgf("initialize http server err=%v", err)
	}

	go func() {
		if err := httpServer.Start(); err != nil {
			log.Fatal().Msgf("start http server got ERR=%v", err)
		}
	}()

	server.GracefulShutdown(httpServer.Stop)
}
