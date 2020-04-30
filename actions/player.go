package actions

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

func addPlayer(name string) (Game, PlayerState, error) {
	if len(currentDeck) < 3 {
		return currentGame, PlayerState{}, errors.New("Deck not big enough to make a new hand")
	}

	if len(currentGame.PlayerStates) == 4 {
		return currentGame, PlayerState{}, errors.New("No more new players can be added")
	}

	uuid, _ := uuid.NewV4()

	// For testing
	var playerId string
	if debugOn {
		playerId = "player_" + uuid.String()
	} else {
		playerId = "player_" + name
	}

	newPlayer := PlayerState{
		Info: PlayerInfo{
			Name: name,
			ID:   playerId,
		},
		Hand: []Card{
			getAndRemoveRandomCard(&currentDeck),
			getAndRemoveRandomCard(&currentDeck),
			getAndRemoveRandomCard(&currentDeck),
		},
	}

	currentGame = Game{
		Board:        currentGame.Board,
		PlayerStates: append(currentGame.PlayerStates, newPlayer),
	}

	return currentGame, newPlayer, nil
}

func GetPlayerState(playerStates *[]PlayerState, playerId string) (*PlayerState, Game, error) {
	_, player, err := findPlayer(playerStates, playerId)
	return player, currentGame, err
}

func findPlayer(playerStates *[]PlayerState, playerId string) (int, *PlayerState, error) {
	for i, v := range *playerStates {
		if v.Info.ID == playerId {
			return i, &v, nil
		}
	}
	return -1, &PlayerState{}, errors.New("Can't find player")
}
