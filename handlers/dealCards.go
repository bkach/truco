package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

type DealCardsResponse struct {
	Game truco.Game `json:"game"`
}

func DealCardsHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries util.QueryExtractor) {
		gameId, err := queries.Query("game_id")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = truco.DealCards(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		_, game, err := truco.FindGameWithId(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(DealCardsResponse{
			Game: *game,
		})

		fmt.Printf("\nDealt cards to game %s", gameId)

	})
}
