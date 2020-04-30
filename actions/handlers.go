package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
)

type AddPlayerRequest struct {
	Name string
}

type AddPlayerResponse struct {
	PlayerState PlayerState
	Board       Cards
}

type GetPlayerStateRequest struct {
	PlayerId string
}

type GetPlayerStateResponse struct {
	PlayerState PlayerState
	Board       Cards
}

type PlayCardRequest struct {
	Card Card
	ID   string
}

type PlayCardResponse struct {
	PlayerState PlayerState
	Board       Cards
}

// Handles requests to start the game
func NewGameHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(startGame()))
}

// Handles requests to add a new player
func AddPlayerHandler(c buffalo.Context) error {
	name := &AddPlayerRequest{}
	c.Bind(name)

	game, playerState, err := addPlayer(name.Name)

	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(
			map[string]string{"error": err.Error()},
		))
	}

	return c.Render(http.StatusOK, r.JSON(AddPlayerResponse{
		PlayerState: playerState,
		Board:       game.Board,
	}))
}

// Handles requests to get a player state
func GetPlayerStateHandler(c buffalo.Context) error {
	request := &GetPlayerStateRequest{}
	c.Bind(request)

	playerState, game, err := currentGame.PlayerStates.GetPlayerState(request.PlayerId)

	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(
			map[string]string{"error": err.Error()},
		))
	}

	return c.Render(http.StatusOK, r.JSON(AddPlayerResponse{
		PlayerState: *playerState,
		Board:       game.Board,
	}))
}

// Handles requests to play a card
func PlayCardHandler(c buffalo.Context) error {
	request := &PlayCardRequest{}
	c.Bind(request)

	playerState, game, _ := playCard(request.Card, request.ID)

	return c.Render(http.StatusOK, r.JSON(
		PlayCardResponse{
			PlayerState: playerState,
			Board:       game.Board,
		},
	))
}
