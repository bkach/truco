package game

import (
	"errors"
	"github.com/gofrs/uuid"
)

type PlayerInfo struct {
	Name string
	ID   string
}

type PlayerState struct {
	Info PlayerInfo
	Hand []Card
}

func createPlayer(game *Game, name string) (*PlayerState, error) {
	playerUUID, err := uuid.NewV4()

	if err != nil {
		return nil, errors.New("error making UUID")
	}

	// For testing
	var playerId string
	if debugOn {
		playerId = "player_" + name
	} else {
		playerId = "player_" + playerUUID.String()
	}

	var hand []Card
	for i := 0; i < NumCardsInHand; i++ {
		hand = append(hand, popRandomCard(&game.deck))
	}

	return &PlayerState{
		Info: PlayerInfo{
			Name: name,
			ID:   playerId,
		},
		Hand: hand,
	}, nil
}

func FindPlayer(games []Game, gameId string, playerId string) (int, *PlayerState, error) {
	gameIndex, err := FindGameIndex(games, gameId)

	if err != nil {
		return 0, nil, errors.New("cannot find game " + gameId)
	}

	var playerNames []string
	playerStates := games[gameIndex].Players

	for i, v := range playerStates {
		if v.Info.ID == playerId {
			playerNames = append(playerNames, v.Info.ID)
			return i, &v, nil
		}
	}

	return -1, nil, errors.New("can't find player " + playerId + " in game " + games[gameIndex].ID)
}
