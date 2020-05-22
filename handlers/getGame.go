package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

func GetGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, r *http.Request) {
		gameId, err := util.GetQuery(r, "game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		_, selectedGame, err := truco.FindGameWithId(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(selectedGame)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		fmt.Printf("\nGetting game %s", gameId)
	})
}
