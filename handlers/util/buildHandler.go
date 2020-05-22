package util

import (
	"errors"
	"net/http"
)

// Builds an http.HandlerFunc which decodes the json from the request into the dst struct
// and subsequently performs the action() function
func BuildHandler(
	requestBody interface{},
	action http.HandlerFunc,
) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if requestBody != nil {
			err := decodeJSONBody(w, r, requestBody)
			if err != nil {
				var malformedRequest *malformedRequest
				if errors.As(err, &malformedRequest) {
					http.Error(w, malformedRequest.msg, malformedRequest.status)
				} else {
					LogInternalError(w, err)
				}
				return
			}
		}

		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Add("Content-Type", "application/json")

		action(w, r)
	}
}
