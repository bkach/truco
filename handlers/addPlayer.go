package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"truco-backend/handlers/util"
	"truco-backend/truco"
)

type AddPlayerRequestBody struct {
	GameId string `json:"game_id"`
	Name   string `json:"name"`
}

type AddPlayerResponse struct {
	PlayerId	string `json:"player_id"`
}

func AddPlayerHandler() http.HandlerFunc {
	requestBody := AddPlayerRequestBody{}
	return util.BuildHandler(&requestBody, func(w http.ResponseWriter, queries url.Values) {
		playerId, err := truco.AddPlayer(requestBody.GameId, requestBody.Name)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(AddPlayerResponse{
			PlayerId: playerId,
		})
	})
}
