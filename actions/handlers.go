package actions

import (
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

type GetGameRequest struct {
	GameId string
}

type GetGameResponse struct {
	Board   []game.Card
	Players []PlayerResponse
}

type PlayerResponse struct {
	Name string
	Hand []game.Card
	ID   string
}

type DealCardsRequest struct {
	GameId string
}

// Handles requests to start the game
func createGameHandler(c buffalo.Context) error {
	gameId, err := game.CreateGameAndAddToGames()

	if err != nil {
		return renderError(c, err)
	}

	return c.Render(http.StatusOK, r.JSON(CreateGameResponse{GameId: gameId}))
}

// Handles requests to get a list of game ids
func getGameIdsHandler(c buffalo.Context) error {
	return c.Render(http.StatusOK, r.JSON(GetGameIdsResponse{
		GameIds: game.GetGameIds(game.Games),
	}))
}

// Handles requests to get a list of game ids
func getGameHandler(c buffalo.Context) error {
	request := &GetGameRequest{}
	err := c.Bind(request)

	if err != nil {
		return renderError(c, err)
	}

	selectedGame := game.Games[request.GameId]

	var playerResponse []PlayerResponse
	for key, playerState := range selectedGame.Players {
		playerResponse = append(playerResponse, PlayerResponse{
			Name: playerState.Name,
			Hand: playerState.Hand,
			ID:   key,
		})
	}

	return c.Render(http.StatusOK, r.JSON(GetGameResponse{
		Board:   selectedGame.Board,
		Players: playerResponse,
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

	return c.Render(http.StatusOK, nil)
}

// Handles requests to get a player state
func getPlayerStateHandler(c buffalo.Context) error {
	request := &GetPlayerStateRequest{}
	err := c.Bind(request)

	if err != nil {
		return renderError(c, err)
	}

	return c.Render(http.StatusOK, r.JSON(GetPlayerStateResponse{
		PlayerState: game.Games[request.GameId].Players[request.PlayerId],
		Board:       game.Games[request.GameId].Board,
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

	return c.Render(http.StatusOK, nil)
}

func dealCardsHandler(c buffalo.Context) error {
	request := &DealCardsRequest{}
	err := c.Bind(request)

	if err != nil {
		return renderError(c, err)
	}

	err = game.DealCards(game.Games, request.GameId)

	if err != nil {
		return renderError(c, err)
	}

	return c.Render(http.StatusOK, nil)
}

func renderError(c buffalo.Context, err error) error {
	return c.Render(http.StatusInternalServerError, r.JSON(
		map[string]string{"message": err.Error()},
	))
}
