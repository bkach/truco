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

func CreatePlayer(name string) (*PlayerState, error) {
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
		hand = append(hand, popRandomCard(&CurrentGame.Deck))
	}

	return &PlayerState{
		Info: PlayerInfo{
			Name: name,
			ID:   playerId,
		},
		Hand: hand,
	}, nil
}

func FindPlayer(playerStates []PlayerState, playerId string) (int, *PlayerState, error) {
	var playerNames []string

	for i, v := range playerStates {
		if v.Info.ID == playerId {
			playerNames = append(playerNames, v.Info.ID)
			return i, &v, nil
		}
	}

	return -1, nil, errors.New("can't find player")
}
