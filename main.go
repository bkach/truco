package main

import (
	"log"
	"net/http"
)

func main() {
	router := setupRouter()
	log.Fatal(http.ListenAndServe(":8000", router))
}
