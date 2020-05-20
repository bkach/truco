package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type CreatePlayerResponse struct {
	Player truco.Player `json:"player"`
	Game truco.Game `json:"game"`
}

func CreatePlayerHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries util.QueryExtractor) {
		gameId, err := queries.Query("game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

	    _, game, err := truco.FindGameWithId(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		playerName, err := queries.Query("name")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		player, err := truco.CreatePlayer(gameId, playerName)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(CreatePlayerResponse{
			Player: *player,
			Game: *game,
		})

		fmt.Printf("\nCreated game %s with playerId %s", playerName, player.Id)
	})
}
