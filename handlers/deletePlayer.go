package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type DeletePlayerRequest struct {
	GameId string `json:"game_id"`
}

type DeletePlayerResponse struct {
	Game truco.Game `json:"game"`
}

func DeletePlayerHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries util.QueryExtractor) {
		gameId, err := queries.Query("game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		playerId, err := queries.Query("player_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = truco.DeletePlayer(gameId, playerId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		_, game, err := truco.FindGameWithId(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(DeletePlayerResponse{
			Game: *game,
		})

		fmt.Printf("\nDeleted game with gameId %s", gameId)
	})
}
