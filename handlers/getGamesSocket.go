package handlers

import (
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
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

func GetGamesSocketHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Upgrades the http session to a websocket session
		conn, err := upgrader.Upgrade(w, r, nil)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		// Only close the connection when the surrounding function returns
		defer func() {
			// Remove the game change listener associated with this address
			truco.RemoveGameChangeListener(conn.RemoteAddr().String())

			// Then close the connection
			err := conn.Close()

			if err != nil {
				util.LogInternalError(w, err)
				return
			}
		}()

		// Create a game change listener for this connection
		var gameChangeListener = func() {
			err := conn.WriteJSON(truco.Games)

			if err != nil {
				util.LogInternalError(w, err)
				return
			}
		}

		// Register the game change listener to the connection's ID
		truco.RegisterGameChangeListener(conn.RemoteAddr().String(), gameChangeListener)

		// Listen for messages - not really used for any functionality at the moment
		for {
			messageType, p, err := conn.ReadMessage()

			if err != nil {
				log.Println(err.Error())
				return
			}

			if messageType == websocket.TextMessage {
				fmt.Printf("Socket: Message Recieved: %s\n", string(p))
			} else if messageType == websocket.CloseMessage {
				fmt.Printf("Socket connection closed %+v\n", p)
			}
		}
	}
}
