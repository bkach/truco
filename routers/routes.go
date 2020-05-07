package routers

import (
	"github.com/gorilla/mux"
	"truco-backend/handlers"
)

func BuildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	// The route below is currently just used for testing
	router.HandleFunc("/", handlers.CreateGameHandler())

	// Creates a game
	router.HandleFunc("/createGame", handlers.CreateGameHandler())

	// Gets a game, game_id query required
	router.HandleFunc("/game", handlers.GetGameHandler())

	// Gets games
	router.HandleFunc("/games", handlers.GetGamesHandler())

	// Adds a player
	router.HandleFunc("/addPlayer", handlers.AddPlayerHandler())
	return router
}
