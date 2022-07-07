package test

import (
	"github.com/gin-gonic/gin"
	"hyuki-go-app/services/websocket"
)

type EchoController struct{}

func (h EchoController) Echo(c *gin.Context) {
	handler := websocket.NewHandler()
	handler.Echo(c.Writer, c.Request)
}
