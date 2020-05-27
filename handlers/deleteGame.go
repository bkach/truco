package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

func DeleteGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, r *http.Request) {
		gameId, err := util.GetQuery(r, "game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = truco.DeleteGame(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(nil)

		fmt.Printf("\nDeleted game with gameId %s", gameId)
	})
}
