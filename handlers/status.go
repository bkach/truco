package handlers

import (
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

// Handler which ...
func StatusHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries util.QueryExtractor) {
		numGames := truco.NumGames()

		var statusString string
		if numGames == 1 {
			statusString = "Truco up and running, %d active game"
		} else {
			statusString = "Truco up and running, %d active games"
		}
		_, err := fmt.Fprintf(w, statusString, numGames)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}
	})
}
