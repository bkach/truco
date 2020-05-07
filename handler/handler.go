package handler

import (
	"errors"
	"log"
	"net/http"
)

func logInternalError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

// Builds an http.HandlerFunc which decodes the json from the request into the dst struct
// and subsequently performs the action() function
func buildHandler(dst interface{}, action func(w http.ResponseWriter)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if dst != nil {
			err := decodeJSONBody(w, r, dst)
			if err != nil {
				var malformedRequest *malformedRequest
				if errors.As(err, &malformedRequest) {
					http.Error(w, malformedRequest.msg, malformedRequest.status)
				} else {
					logInternalError(w, err)
				}
				return
			}
		}

		action(w)
	}
}
