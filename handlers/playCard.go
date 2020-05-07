package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"truco-backend/handlers/util"
	"truco-backend/truco"
)

type PlayCardRequest struct {
	Card truco.Card `json:"card"`
}

func PlayCardHandler() http.HandlerFunc {
	request := PlayCardRequest{}
	return util.BuildHandler(&request, func(w http.ResponseWriter, queries url.Values) {
		gameId := queries["game_id"][0]
		playerId := queries["player_id"][0]

		if gameId == "" || playerId == "" {
			errorString := fmt.Sprintf("cannot find user with gameId \"%s\" and player id \"%s\"", gameId, playerId)
			util.LogInternalError(w, errors.New(errorString))
			return
		}

		err := truco.PlayCard(gameId, playerId, request.Card)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
