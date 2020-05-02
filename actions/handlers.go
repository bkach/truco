package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/truco_backend/game"
	"net/http"
)

type CreateGameResponse struct {
	GameId string `json:"game_id"`
}

type AddPlayerRequest struct {
	GameId string `json:"game_id"`
	Name   string `json:"name"`
}

type GetPlayerStateRequest struct {
	GameId   string `json:"game_id"`
	PlayerId string `json:"player_id"`
}

type GetPlayerStateResponse struct {
	PlayerResponse PlayerResponse `json:"player_response"`
	Board          []game.Card    `json:"board"`
}

type PlayCardRequest struct {
	GameId   string    `json:"game_id"`
	PlayerId string    `json:"player_id"`
	Card     game.Card `json:"card"`
}

type GetGameIdsResponse struct {
	GameIds []string `json:"game_ids"`
}

type GetGameRequest struct {
	GameId string `json:"game_id"`
}

type GetGameResponse struct {
	Board   []game.Card      `json:"board"`
	Players []PlayerResponse `json:"players"`
}

type PlayerResponse struct {
	Name string      `json:"name"`
	Hand []game.Card `json:"hand"`
	ID   string      `json:"id"`
}

type DealCardsRequest struct {
	GameId string `json:"game_id"`
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

	player := game.Games[request.GameId].Players[request.PlayerId]

	var playerResponse = PlayerResponse{
		Name: player.Name,
		Hand: player.Hand,
		ID:   request.PlayerId,
	}

	return c.Render(http.StatusOK, r.JSON(GetPlayerStateResponse{
		PlayerResponse: playerResponse,
		Board:          game.Games[request.GameId].Board,
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
