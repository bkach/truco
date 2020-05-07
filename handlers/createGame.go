package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"truco-backend/handlers/util"
	"truco-backend/truco"
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
