package handlers

import (
	"encoding/json"
	"net/http"
	"net/url"
	"truco-backend/handlers/util"
	"truco-backend/truco"
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
