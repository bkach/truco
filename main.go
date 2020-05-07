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
	router.HandlerFunc("/games", handler.GetGamesHandler())

	//router.HandlerFunc("/dealCards", handler.DealCardsHandler)
	//router.HandlerFunc("/addPlayer", handler.AddPlayerHandler)
	//router.HandlerFunc("/playCard", handler.PlayCardHandler)

	log.Fatal(http.ListenAndServe(":8000", router))
}
