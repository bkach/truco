package truco

import (
	"github.com/gofrs/uuid"
	"strconv"
)

type Game struct {
	Name    string   `json:"name"`
	Id      string   `json:"id"`
	Players []Player `json:"players"`
	deck    []Card
}

func createGame(games []Game, name string) (*Game, error) {
	gameUUID, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	// For testing
	var gameId string
	if debugOn {
		gameId = "game_" + strconv.Itoa(len(games))
	} else {
		gameId = "game_" + gameUUID.String()
	}

	newGame := Game{
		Name:    name,
		Id:      gameId,
		Players: []Player{},
		deck:    buildDeck(),
	}

	return &newGame, nil
}
