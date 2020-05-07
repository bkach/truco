package util

import (
	"errors"
	"net/url"
)

type QueryExtractor struct {
	values url.Values
}

func (extractor *QueryExtractor) Query(query string) (string, error) {
	valueList := extractor.values[query]

	if len(valueList) == 0 {
		return "", errors.New("Query not found: missing \"" + query + "\" query in the URL")
	}

	return valueList[0], nil
}
