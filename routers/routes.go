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

	// Adds a player, query needs a game_id, body needs a name
	router.HandleFunc("/addPlayer", handlers.AddPlayerHandler())

	// Deals in each player for a given game, query needs a game_id
	router.HandleFunc("/dealCards", handlers.DealCardsHandler())

	// Plays a card in a given player's hand
	router.HandleFunc("/playCard", handlers.PlayCardHandler())

	return router
}
