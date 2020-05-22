package sockets

import (
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"github.com/gorilla/websocket"
	"net/http"
)

type PlayerListResponse struct {
	Player []truco.Player `json:"players"`
}

func GetPlayerListSocketHandler() http.HandlerFunc {
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

		truco.RegisterPlayerListChangeListener(
			conn.RemoteAddr().String(),
			gameId,
			playerListChangeListener(gameId, conn),
		)
	}

	onDisconnect := func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		truco.RemovePlayerListChangeListener(conn.RemoteAddr().String())
	}

	return util.BuildSocketHandler(onConnect, onDisconnect)
}

func playerListChangeListener(gameId string, conn *websocket.Conn) func() {
	return func() {
		_, game, err := truco.FindGameWithId(gameId)

		if err != nil {
			util.LogSocketErrorAndCloseConnection(conn, err)
			return
		}

		err = conn.WriteJSON(game.Players)

		if err != nil {
			util.LogSocketErrorAndCloseConnection(conn, err)
			return
		}
	}
}

