package handler

import (
	"encoding/json"
	"net/http"
	"truco-backend/truco"
)

type GetGameRequest struct {
	GameId string `json:"game_id"`
}

// Handler which gets a specific game of truco
func GetGameHandler() http.HandlerFunc {
	request := GetGameRequest{}
	return buildHandler(&request, func(w http.ResponseWriter) {
		_, selectedGame, err := truco.FindGameWithId(request.GameId)

		if err != nil {
			logInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(selectedGame)

		if err != nil {
			logInternalError(w, err)
			return
		}
	})
}
