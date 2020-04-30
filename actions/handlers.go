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
	bindErr := c.Bind(name)

	if bindErr != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": bindErr.Error()},
		))
	}

	playerState, addErr := game.AddPlayer(name.Name)

	if addErr != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": addErr.Error()},
		))
	}

	if playerState == nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": "player could not be created"},
		))
	}

	return c.Render(http.StatusOK, r.JSON(AddPlayerResponse{
		PlayerState: *playerState,
		Board:       game.CurrentGame.Board,
	}))
}

// Handles requests to get a player state
func GetPlayerStateHandler(c buffalo.Context) error {
	request := &GetPlayerStateRequest{}
	err := c.Bind(request)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": err.Error()},
		))
	}

	_, playerState, err := game.FindPlayer(game.CurrentGame.Players, request.PlayerId)

	if err != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": err.Error()},
		))
	}

	return c.Render(http.StatusOK, r.JSON(AddPlayerResponse{
		PlayerState: *playerState,
		Board:       game.CurrentGame.Board,
	}))
}

// Handles requests to play a card
func PlayCardHandler(c buffalo.Context) error {
	request := &PlayCardRequest{}
	errOnBind := c.Bind(request)

	if errOnBind != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": errOnBind.Error()},
		))
	}

	errOnPlay := game.PlayCard(request.Card, request.ID)

	if errOnPlay != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": errOnPlay.Error()},
		))
	}

	_, requestedPlayerState, errOnFind := game.FindPlayer(game.CurrentGame.Players, request.ID)

	if errOnFind != nil {
		return c.Render(http.StatusInternalServerError, r.JSON(
			map[string]string{"message": errOnFind.Error()},
		))
	}

	return c.Render(http.StatusOK, r.JSON(
		PlayCardResponse{
			PlayerState: *requestedPlayerState,
			Board:       game.CurrentGame.Board,
		},
	))
}
