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

func AddPlayer(name string) (PlayerState, Game, error) {
	if len(currentDeck) < 3 {
		return PlayerState{}, CurrentGame, errors.New("deck not big enough to make a new hand")
	}

	if len(CurrentGame.PlayerStates) == 4 {
		return PlayerState{}, CurrentGame, errors.New("no more new players can be added")
	}

	playerUUID, err := uuid.NewV4()

	if err != nil {
		return PlayerState{}, CurrentGame, errors.New("error making UUID")
	}

	// For testing
	var playerId string
	if debugOn {
		playerId = "player_" + playerUUID.String()
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

	CurrentGame = Game{
		Board:        CurrentGame.Board,
		PlayerStates: append(CurrentGame.PlayerStates, newPlayer),
	}

	return newPlayer, CurrentGame, nil
}

func GetPlayerState(playerStates *[]PlayerState, playerId string) (*PlayerState, Game, error) {
	_, player, err := findPlayer(playerStates, playerId)
	return player, CurrentGame, err
}

func findPlayer(playerStates *[]PlayerState, playerId string) (int, *PlayerState, error) {
	for i, v := range *playerStates {
		if v.Info.ID == playerId {
			return i, &v, nil
		}
	}
	return -1, &PlayerState{}, errors.New("Can't find player")
}
