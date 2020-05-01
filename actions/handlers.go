package actions

import (
	"fmt"
	"github.com/gobuffalo/buffalo"
	"github.com/truco_backend/game"
	"net/http"
)

type CreateGameResponse struct {
	GameId string
}

type AddPlayerRequest struct {
	GameId string
	Name   string
}

type GetPlayerStateRequest struct {
	GameId   string
	PlayerId string
}

type GetPlayerStateResponse struct {
	PlayerState game.PlayerState
	Board       []game.Card
}

type PlayCardRequest struct {
	GameId   string
	PlayerId string
	Card     game.Card
}

type GetGameIdsResponse struct {
	GameIds []string
}

// Handles requests to start the game
func createGameHandler(c buffalo.Context) error {
	fmt.Print("\n\nIn the handler\n\n")
	gameId, err := game.CreateGameAndAddToGames()

	if err != nil {
		return renderError(c, err)
	}

	fmt.Printf("%v", game.Games)
	return c.Render(http.StatusOK, r.JSON(CreateGameResponse{GameId: gameId}))
}

// Handles requests to get a list of game ids
func getGameIdsHandler(c buffalo.Context) error {
	fmt.Printf("%v", game.Games)
	return c.Render(http.StatusOK, r.JSON(GetGameIdsResponse{
		GameIds: game.GetGameIds(game.Games),
	}))
}

// Handles requests to add a new player
func addPlayerHandler(c buffalo.Context) error {
	request := &AddPlayerRequest{}
	err := c.Bind(request)

	if err != nil {
		return renderError(c, err)
	}

	err = game.AddPlayer(game.Games, request.GameId, request.Name)

	if err != nil {
		return renderError(c, err)
	}

	fmt.Printf("%v", game.Games)
	return c.Render(http.StatusOK, nil)
}

// Handles requests to get a player state
func getPlayerStateHandler(c buffalo.Context) error {
	request := &GetPlayerStateRequest{}
	err := c.Bind(request)

	if err != nil {
		return renderError(c, err)
	}

	gameIndex, err := game.FindGameIndex(game.Games, request.GameId)

	if err != nil {
		return renderError(c, err)
	}

	_, playerState, err := game.FindPlayer(game.Games, request.GameId, request.PlayerId)

	if err != nil {
		return renderError(c, err)
	}

	fmt.Printf("%v", game.Games)
	return c.Render(http.StatusOK, r.JSON(GetPlayerStateResponse{
		PlayerState: *playerState,
		Board:       game.Games[gameIndex].Board,
	}))
}

// Handles requests to play a card
func playCardHandler(c buffalo.Context) error {
	request := &PlayCardRequest{}
	err := c.Bind(request)

	if err != nil {
		return renderError(c, err)
	}

	err = game.PlayCard(game.Games, request.Card, request.GameId, request.PlayerId)

	fmt.Printf("%v", game.Games)
	return c.Render(http.StatusOK, nil)
}

func renderError(c buffalo.Context, err error) error {
	return c.Render(http.StatusInternalServerError, r.JSON(
		map[string]string{"message": err.Error()},
	))
}
