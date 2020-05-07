package handler

import (
	"encoding/json"
	"net/http"
	"truco-backend/truco"
)

// Handler which ...
func GetGamesHandler() http.HandlerFunc {
	return buildHandler(nil, func(w http.ResponseWriter) {
		err := json.NewEncoder(w).Encode(truco.Games)

		if err != nil {
			logInternalError(w, err)
			return
		}
	})
}
