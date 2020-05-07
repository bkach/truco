package handlers

import (
	"github.com/bkach/truco-backend/handlers/util"
	"github.com/bkach/truco-backend/truco"
	"net/http"
	"net/url"
)

func DeleteGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries url.Values) {
		gameId := queries["game_id"][0]

		err := truco.DeleteGame(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
