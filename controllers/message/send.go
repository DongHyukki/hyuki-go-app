package message

import (
	"github.com/gin-gonic/gin"
	"hyuki-go-app/services/websocket"
	"io/ioutil"
	"log"
)

type sendController struct {
	handler *websocket.Handler
}

func NewController() *sendController {
	return &sendController{handler: websocket.NewHandler()}
}

func (m sendController) Send(c *gin.Context) {
	sendMessage, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		return
	}

	log.Printf("Read Bytes Size :: %d", len(sendMessage))

	m.handler.WriteMessage(c.Request.Context(), string(sendMessage))
}
