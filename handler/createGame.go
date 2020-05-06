package handler

import (
	"encoding/json"
	"net/http"
	"truco-backend/truco"
)

type CreateGameResponse struct {
	GameId string `json:"game_id"`
}

// Handler which creates a game
func CreateGameHandler() http.HandlerFunc {
	return buildHandler(nil, func(w http.ResponseWriter) {
		gameId, err := truco.CreateGameAndAddToGames()

		if err != nil {
			logInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(CreateGameResponse{
			GameId: gameId,
		})

		if err != nil {
			logInternalError(w, err)
			return
		}
	})
}