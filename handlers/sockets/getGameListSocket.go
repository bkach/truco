package sockets

import (
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"github.com/gorilla/websocket"
	"net/http"
)

type GameListResponse struct {
	Games []GameListResponseGame `json:"games"`
}

type GameListResponseGame struct {
	GameId   string `json:"game_id"`
	GameName string `json:"name"`
}

func GetGameListSocketHandler() http.HandlerFunc {
	onConnect := func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		truco.RegisterGameListChangeListener(
			conn.RemoteAddr().String(),
			gameListChangeListener(conn),
		)
	}

	onDisconnect := func(w http.ResponseWriter, r *http.Request, conn *websocket.Conn) {
		truco.RemoveGameListChangeListener(conn.RemoteAddr().String())
	}

	return util.BuildSocketHandler(onConnect, onDisconnect)
}

func gameListChangeListener(conn *websocket.Conn) func() {
	return func() {
		var responseGames []GameListResponseGame
		for _, game := range truco.Games {
			responseGames = append(responseGames, GameListResponseGame{
				GameId:   game.Id,
				GameName: game.Name,
			})
		}

		response := GameListResponse{
			Games: responseGames,
		}

		err := conn.WriteJSON(response)

		if err != nil {
			util.LogSocketErrorAndCloseConnection(conn, err)
			return
		}
	}
}

