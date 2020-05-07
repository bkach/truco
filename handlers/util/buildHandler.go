package util

import (
	"errors"
	"net/http"
	"net/url"
)

// Builds an http.HandlerFunc which decodes the json from the request into the dst struct
// and subsequently performs the action() function
func BuildHandler(
	requestBody interface{},
	action func(w http.ResponseWriter, queryValues url.Values),
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

		action(w, r.URL.Query())
	}
}
