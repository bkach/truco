package util

import (
	"errors"
	"net/http"
)

func GetQuery(r *http.Request, query string) (string, error) {
	valueList := r.URL.Query()[query]

	if len(valueList) == 0 {
		return "", errors.New("Query not found: missing \"" + query + "\" query in the URL")
	}

	return valueList[0], nil
}
