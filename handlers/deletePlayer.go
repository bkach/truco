package handlers

import (
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type DeletePlayerRequest struct {
	GameId string `json:"game_id"`
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

		w.WriteHeader(http.StatusOK)

		fmt.Printf("\nDeleted game with gameId %s", gameId)
	})
}
