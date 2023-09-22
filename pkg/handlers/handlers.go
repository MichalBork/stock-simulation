package handlers

import (
	"github.com/gorilla/websocket"
	"net/http"
)

func NewWebsocket(conn *websocket.Conn) {
	// Handle your websocket connection here
	// For example, you can start a new goroutine for each connection to read
	// messages:

	go func() {
		for {
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				return
			}
			if err := conn.WriteMessage(messageType, p); err != nil {
				return
			}
		}
	}()
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	// Handle your HTTP request here
}
