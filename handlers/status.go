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
		_, err := fmt.Fprintf(w, "Truco up and running, %d games", truco.NumGames())

		if err != nil {
			util.LogInternalError(w, err)
			return
		}
	})
}
