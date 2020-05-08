package routers

import (
	"github.com/bkach/truco/handlers"
	"github.com/gorilla/mux"
)

func BuildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// The route below is currently just used for testing
	router.HandleFunc("/", handlers.StatusHandler())

	// Creates a game
	router.HandleFunc("/createGame", handlers.CreateGameHandler())

	// Deletes a game, game_id query required
	router.HandleFunc("/deleteGame", handlers.DeleteGameHandler())

	// Gets a game, game_id query required
	router.HandleFunc("/game", handlers.GetGameHandler())

	// Gets games
	router.HandleFunc("/games", handlers.GetGamesHandler())

	// Creates a player, query needs a game_id and a name for the player
	router.HandleFunc("/createPlayer", handlers.CreatePlayerHandler())

	// Removes a player, query needs a game_id and a player_id
	router.HandleFunc("/deletePlayer", handlers.DeletePlayerHandler())

	// Deals in each player for a given game, query needs a game_id
	router.HandleFunc("/dealCards", handlers.DealCardsHandler())

	// Plays a card in a given player's hand, query needs a game_id and body needs a Card
	router.HandleFunc("/playCard", handlers.PlayCardHandler())

	return router
}
