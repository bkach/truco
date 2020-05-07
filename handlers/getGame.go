package handlers

import (
	"encoding/json"
	"github.com/bkach/truco-backend/handlers/util"
	"github.com/bkach/truco-backend/truco"
	"net/http"
	"net/url"
)

func GetGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries url.Values) {
		gameId := queries["game_id"][0]

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
	})
}
