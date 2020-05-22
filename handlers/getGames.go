package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/bkach/truco/handlers/util"
	"github.com/bkach/truco/truco"
	"net/http"
)

func GetGamesHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, r *http.Request) {
		err := json.NewEncoder(w).Encode(truco.Games)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		fmt.Print("\nGetting all games")
	})
}
