package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type PlayCardRequest struct {
	Card truco.Card `json:"card"`
}

type PlayCardResponse  struct {
	Game truco.Game `json:"game"`
}

func PlayCardHandler() http.HandlerFunc {
	request := PlayCardRequest{}
	return util.BuildHandler(&request, func(w http.ResponseWriter, queries util.QueryExtractor) {
		gameId, err := queries.Query("game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		playerId, err := queries.Query("player_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = truco.PlayCard(gameId, playerId, request.Card)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		_, game, err := truco.FindGameWithId(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(PlayCardResponse{
			Game: *game,
		})

		fmt.Printf("\nPlaying card %+v %s", request.Card, gameId)
	})
}
