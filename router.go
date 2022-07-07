package main

import (
	"github.com/gin-gonic/gin"
	"hyuki-go-app/controllers/test"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	healthController := new(test.HealthController)
	echoController := new(test.EchoController)

	v1 := r.Group("/api")
	{
		v1.GET("/test", func(c *gin.Context) {
			healthController.Status(c)
		})

		websocket := v1.Group("/websocket")
		websocket.GET("/echo", func(c *gin.Context) {
			echoController.Echo(c)
		})
	}

	return r
}
