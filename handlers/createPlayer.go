package handlers

import (
	"encoding/json"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type CreatePlayerResponse struct {
	PlayerId string `json:"player_id"`
}

func CreatePlayerHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries util.QueryExtractor) {
		gameId, err := queries.Query("game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		playerName, err := queries.Query("name")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		playerId, err := truco.CreatePlayer(gameId, playerName)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(CreatePlayerResponse{
			PlayerId: playerId,
		})
	})
}
