package websocket

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

type Handler struct {
	upgrader websocket.Upgrader
}

func NewHandler() *Handler {
	return &Handler{
		upgrader: websocket.Upgrader{},
	}
}

func (h Handler) Echo(w http.ResponseWriter, r *http.Request) {
	c, err := h.upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}
		log.Printf("recv: %s", message)
		err = c.WriteMessage(mt, message)
		if err != nil {
			log.Println("write:", err)
			break
		}
	}
}