package truco

import (
	"github.com/gofrs/uuid"
	"strconv"
)

type Game struct {
	Id      string                 `json:"id"`
	Players map[string]PlayerState `json:"players"`
	Board   []Card                 `json:"board"`
	Deck    []Card                 `json:"deck"`
}

func createGame(games []Game) (*Game, error) {
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
		Id:      gameId,
		Board:   []Card{},
		Players: map[string]PlayerState{},
		Deck:    buildDeck(),
	}

	return &newGame, nil
}
