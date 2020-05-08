package handlers

import (
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

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

		w.WriteHeader(http.StatusOK)

		fmt.Printf("\nDealt cards to game %s", gameId)

	})
}
