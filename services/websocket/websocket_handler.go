package websocket

import (
	"context"
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
		break
	}
}

func (h Handler) WriteMessage(ctx context.Context, msg string) {
	c, _, err := websocket.DefaultDialer.DialContext(ctx, "ws://localhost:8080/api/websocket/echo", nil)
	if err != nil {
		log.Fatal("dial:", err)
	}
	defer c.Close()

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			_, message, err := c.ReadMessage()
			if err != nil {
				if websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Println("read finish from echo")
					return
				}
				log.Printf("error: %v", err)
				break
			}
			log.Printf("recv from echo: %s", message)
		}
	}()

	err1 := c.WriteMessage(websocket.BinaryMessage, []byte(msg))
	if err1 != nil {
		log.Println("write:", err)
		return
	}

	for {
		select {
		case <-done:
			return
		}
	}
}
