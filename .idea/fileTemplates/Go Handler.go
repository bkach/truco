package ${GO_PACKAGE_NAME}

import (
	"encoding/json"
	"net/http"
	"truco-backend/truco"
)

type ${className}Request struct {
	GameId string `json:"game_id"`
}

// Handler which ...
func ${className}Handler() http.HandlerFunc {
	request := GetGameRequest{}
	return buildHandler(&request, func(w http.ResponseWriter) {
	    err := Errors.New("test error")

		if err != nil {
			logInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(...)

		if err != nil {
			logInternalError(w, err)
			return
		}
	})
}
