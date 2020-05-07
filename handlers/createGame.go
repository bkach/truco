package handlers

import (
	"encoding/json"
	"github.com/bkach/truco-backend/handlers/util"
	"github.com/bkach/truco-backend/truco"
	"net/http"
	"net/url"
)

type CreateGameResponse struct {
	GameId string `json:"game_id"`
}

func CreateGameHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries url.Values) {
		gameId, err := truco.CreateGameAndAddToGames()

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(CreateGameResponse{
			GameId: gameId,
		})

		if err != nil {
			util.LogInternalError(w, err)
			return
		}
	})
}
