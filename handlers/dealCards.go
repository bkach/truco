package handlers

import (
	"errors"
	"fmt"
	"github.com/bkach/truco-backend/handlers/util"
	"github.com/bkach/truco-backend/truco"
	"net/http"
	"net/url"
)

func DealCardsHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries url.Values) {
		gameId := queries["game_id"][0]

		if gameId == "" {
			errorString := fmt.Sprintf("cannot find user with gameId \"%s\" ", gameId)
			util.LogInternalError(w, errors.New(errorString))
			return
		}

		err := truco.DealCards(gameId)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		w.WriteHeader(http.StatusOK)
	})
}
