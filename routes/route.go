package routes

import (
	"github.com/gin-gonic/gin"
	"web-app/logger"
)

func Setup() *gin.Engine {
	r := gin.New()
	r.Use(logger.GinLogger(), logger.GinRecovery(true))
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}
