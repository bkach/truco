package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"truco-backend/handlers/util"
	"truco-backend/truco"
)

type AddPlayerRequestBody struct {
	Name string `json:"name"`
}

type AddPlayerResponse struct {
	PlayerId string `json:"player_id"`
}

func AddPlayerHandler() http.HandlerFunc {
	requestBody := AddPlayerRequestBody{}
	return util.BuildHandler(&requestBody, func(w http.ResponseWriter, queries url.Values) {
		gameId := queries["game_id"][0]
		playerId, err := truco.AddPlayer(gameId, requestBody.Name)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(AddPlayerResponse{
			PlayerId: playerId,
		})
	})
}
