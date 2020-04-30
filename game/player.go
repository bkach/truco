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

func AddPlayer(name string) (*PlayerState, error) {
	if len(CurrentGame.Deck) < NUM_CARDS_IN_HAND {
		return nil, errors.New("deck not big enough to make a new hand")
	}

	if len(CurrentGame.Players) == MAX_PLAYERS {
		return nil, errors.New("no more new players can be added")
	}

	playerUUID, err := uuid.NewV4()

	if err != nil {
		return nil, errors.New("error making UUID")
	}

	// For testing
	var playerId string
	if debugOn {
		playerId = "player_" + playerUUID.String()
	} else {
		playerId = "player_" + name
	}

	var hand []Card
	for i := 0; i < NUM_CARDS_IN_HAND; i++ {
		hand = append(hand, popRandomCard(&CurrentGame.Deck))
	}

	newPlayer := PlayerState{
		Info: PlayerInfo{
			Name: name,
			ID:   playerId,
		},
		Hand: hand,
	}

	CurrentGame = Game{
		Board:   CurrentGame.Board,
		Players: append(CurrentGame.Players, newPlayer),
		Deck:    CurrentGame.Deck,
	}

	return &newPlayer, nil
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
