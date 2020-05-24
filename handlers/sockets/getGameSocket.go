package sockets

import (
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"github.com/gorilla/websocket"
	"net/http"
)

func GetGameSocketHandler() http.HandlerFunc {
	onConnect := func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		gameId, err := util.GetQuery(r, "game_id")

		if err != nil {
			util.LogSocketErrorAndCloseConnection(conn, err)
			return
		}

		_, _, err = truco.FindGameWithId(gameId)

		if err != nil {
			util.LogSocketErrorAndCloseConnection(conn, err)
			return
		}

		truco.RegisterGameChangeListener(
			conn.RemoteAddr().String(),
			gameId,
			gameChangeListener(gameId, conn),
		)
	}

	onDisconnect := func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		truco.RemoveGameChangeListener(conn.RemoteAddr().String())
	}

	return util.BuildSocketHandler(onConnect, onDisconnect)
}

func gameChangeListener(gameId string, conn *websocket.Conn) func() {
	return func() {
		_, game, err := truco.FindGameWithId(gameId)

		if err != nil {
			util.LogSocketErrorAndCloseConnection(conn, err)
			return
		}

		err = conn.WriteJSON(game)

		if err != nil {
			util.LogSocketErrorAndCloseConnection(conn, err)
			return
		}
	}
}

