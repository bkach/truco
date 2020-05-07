package handlers

import (
	"encoding/json"
	"github.com/bkach/truco-backend/handlers/util"
	"github.com/bkach/truco-backend/truco"
	"net/http"
	"net/url"
)

func GetGamesHandler() http.HandlerFunc {
	return util.BuildHandler(nil, func(w http.ResponseWriter, queries url.Values) {
		err := json.NewEncoder(w).Encode(truco.Games)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}
	})
}
