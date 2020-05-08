package routers

import (
	"github.com/bkach/truco/handlers"
	"github.com/gorilla/mux"
	"net/http"
)

func BuildRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	api := router.PathPrefix("/api").Subrouter()

	// Creates a game
	api.HandleFunc("/createGame", handlers.CreateGameHandler())

	// Deletes a game, game_id query required
	api.HandleFunc("/deleteGame", handlers.DeleteGameHandler())

	// Gets a game, game_id query required
	api.HandleFunc("/game", handlers.GetGameHandler())

	// Gets games
	api.HandleFunc("/games", handlers.GetGamesHandler())

	// Creates a player, query needs a game_id and a name for the player
	api.HandleFunc("/createPlayer", handlers.CreatePlayerHandler())

	// Removes a player, query needs a game_id and a player_id
	api.HandleFunc("/deletePlayer", handlers.DeletePlayerHandler())

	// Deals in each player for a given game, query needs a game_id
	api.HandleFunc("/dealCards", handlers.DealCardsHandler())

	// Plays a card in a given player's hand, query needs a game_id and body needs a Card
	api.HandleFunc("/playCard", handlers.PlayCardHandler())

	// The route below serves the frontend
	router.PathPrefix("/").Handler(http.FileServer(http.Dir("frontend/dist/")))

	return router
}
