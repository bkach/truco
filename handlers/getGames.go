package handlers

import (
	"encoding/json"
	"github.com/bkach/truco-backend/handlers/util"
	"github.com/bkach/truco-backend/truco"
	"net/http"
)

func GetGamesHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, _ util.QueryExtractor) {
		err := json.NewEncoder(w).Encode(truco.Games)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}
	})
}
