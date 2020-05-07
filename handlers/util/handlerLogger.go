package util

import (
	"log"
	"net/http"
)

func LogInternalError(w http.ResponseWriter, err error) {
	log.Println(err.Error())
	http.Error(w, "Error: "+err.Error(), http.StatusBadRequest)
}
