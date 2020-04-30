package actions

import (
	"net/http"

	"github.com/gobuffalo/buffalo"
	"github.com/truco_backend/game"
)

type AddPlayerRequest struct {
	Name string
}

type AddPlayerResponse struct {
	PlayerState game.PlayerState
	Board       []game.Card
}

type GetPlayerStateRequest struct {
	PlayerId string
}

type GetPlayerStateResponse struct {
	PlayerState game.PlayerState
	Board       []game.Card
}

type PlayCardRequest struct {
	Card game.Card
	ID   string
}

type PlayCardResponse struct {
	PlayerState game.PlayerState
	Board       []game.Card
}

// Handles requests to start the game
func NewGameHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(game.StartGame()))
}

// Handles requests to add a new player
func AddPlayerHandler(c buffalo.Context) error {
	name := &AddPlayerRequest{}
	err := c.Bind(name)

	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(
			map[string]string{"message": err.Error()},
		))
	}

	playerState, updatedGame, err := game.AddPlayer(name.Name)

	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(
			map[string]string{"message": err.Error()},
		))
	}

	return c.Render(http.StatusOK, r.JSON(AddPlayerResponse{
		PlayerState: playerState,
		Board:       updatedGame.Board,
	}))
}

// Handles requests to get a player state
func GetPlayerStateHandler(c buffalo.Context) error {
	request := &GetPlayerStateRequest{}
	err := c.Bind(request)

	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(
			map[string]string{"message": err.Error()},
		))
	}

	playerState, updatedGame, err := game.GetPlayerState(&game.CurrentGame.PlayerStates, request.PlayerId)

	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(
			map[string]string{"message": err.Error()},
		))
	}

	return c.Render(http.StatusOK, r.JSON(AddPlayerResponse{
		PlayerState: *playerState,
		Board:       updatedGame.Board,
	}))
}

// Handles requests to play a card
func PlayCardHandler(c buffalo.Context) error {
	request := &PlayCardRequest{}
	err := c.Bind(request)

	if err != nil {
		return c.Render(http.StatusBadRequest, r.JSON(
			map[string]string{"message": err.Error()},
		))
	}

	playerState, updatedGame, _ := game.PlayCard(request.Card, request.ID)

	return c.Render(http.StatusOK, r.JSON(
		PlayCardResponse{
			PlayerState: playerState,
			Board:       updatedGame.Board,
		},
	))
}
