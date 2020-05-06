package main

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"truco-backend/handler"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.CreateGameHandler())
	router.HandleFunc("/game", handler.GetGameHandler())

	log.Fatal(http.ListenAndServe(":8000", router))
}
