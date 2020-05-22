package util

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func LogInternalError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
}

type SocketError struct {
	Message string `json:"message"`
}

func LogSocketErrorAndCloseConnection(conn *websocket.Conn, err error) {
	log.Println(err.Error())
	err = conn.WriteJSON(SocketError{
		Message: err.Error() + ", closing socket connection",
	})

	if err != nil {
		log.Println(err.Error())
	}

	err = conn.Close()

	if err != nil {
		log.Println(err.Error())
	}
}
