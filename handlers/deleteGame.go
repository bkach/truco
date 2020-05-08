package handlers

import (
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

func DeleteGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries util.QueryExtractor) {
		gameId, err := queries.Query("game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = truco.DeleteGame(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)

		fmt.Printf("\nDeleted game with gameId %s", gameId)
	})
}
