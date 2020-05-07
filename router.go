package main

import (
	"github.com/gorilla/mux"
	"truco-backend/handler"
)

func setupRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/", handler.CreateGameHandler())
	router.HandleFunc("/game", handler.GetGameHandler())
	router.HandleFunc("/games", handler.GetGamesHandler())
	return router
}