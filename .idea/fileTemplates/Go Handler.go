package ${GO_PACKAGE_NAME}

import (
	"encoding/json"
	"net/http"
	"github.com/bkach/truco-backend/truco"
)

type ${className}Request struct {
	GameId string `json:"game_id"`
}

// Handler which ...
func ${className}Handler() http.HandlerFunc {
	request := ${className}Request{}
	return util.BuildHandler(&request, func(w http.ResponseWriter, queries url.Values) {
	    err := Errors.New("test error")

		if err != nil {
			util.LogInternalError(w, err)
			return
		}

		err = json.NewEncoder(w).Encode(...)

		if err != nil {
			util.LogInternalError(w, err)
			return
		}
	})
}
