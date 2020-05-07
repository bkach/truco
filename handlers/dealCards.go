package handlers

import (
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"truco-backend/handlers/util"
	"truco-backend/truco"
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
