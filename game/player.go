package game

import (
	"errors"
	"github.com/gofrs/uuid"
)

type PlayerState struct {
	Name string
	Hand []Card
}

func createPlayer(name string) (string, *PlayerState, error) {
	playerUUID, err := uuid.NewV4()

	if err != nil {
		return "", nil, errors.New("error making UUID")
	}

	// For testing
	var playerId string
	if debugOn {
		playerId = "player_" + name
	} else {
		playerId = "player_" + playerUUID.String()
	}

	return playerId, &PlayerState{
		Name: name,
	}, nil
}

func dealPlayerIn(deck []Card, player *PlayerState) (*PlayerState, []Card) {
	var hand []Card
	for i := 0; i < NumCardsInHand; i++ {
		hand = append(hand, popRandomCard(&deck))
	}

	player.Hand = hand

	return player, deck
}
