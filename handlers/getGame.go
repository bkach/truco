package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"truco-backend/handlers/util"
	"truco-backend/truco"
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
