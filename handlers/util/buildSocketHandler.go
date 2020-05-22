package util

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

// The upgrader upgrades us from http to a websocket connection. The empty struct means we'll be using
// the default settings
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

type SocketHandlerFunc = func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn)

func BuildSocketHandler(
	onConnect SocketHandlerFunc,
	onDisconnect SocketHandlerFunc,
) http.HandlerFunc {
	return BuildHandler(nil, func(w http.ResponseWriter, r *http.Request) {
		// Upgrades the http session to a websocket session
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			LogInternalError(w, err)
			return
		}

		// Only close the connection when the surrounding function returns
		defer func() {
			onDisconnect(w, r, conn)

			// Then close the connection
			err := conn.Close()

			if err != nil {
				LogInternalError(w, err)
				return
			}
		}()

		onConnect(w, r, conn)

		// Listen for messages - not really used for any functionality at the moment
		for {
			messageType, p, err := conn.ReadMessage()

			if err != nil {
				log.Println(err.Error())
				return
			}

			if messageType == websocket.TextMessage {
				fmt.Printf("\nSocket: Message Recieved: %s", string(p))
			} else if messageType == websocket.CloseMessage {
				fmt.Printf("Socket connection closed %+v\n", p)
			}
		}
	})
}
