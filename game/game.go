package game

import (
	"errors"
)

const MaxPlayers = 4
const NumCardsInHand = 3

type Game struct {
	Board   []Card
	Players []PlayerState
	Deck    []Card
}

// Global state
var CurrentGame Game
var debugOn = false

func StartGame() Game {
	CurrentGame = Game{
		Board:   []Card{},
		Players: []PlayerState{},
		Deck:    buildDeck(),
	}
	return CurrentGame
}

func PlayCard(card Card, playerId string) error {
	playerIndex, _, err := FindPlayer(CurrentGame.Players, playerId)

	if err != nil {
		return err
	}

	playerHand := &CurrentGame.Players[playerIndex].Hand

	index, err := findCardIndex(playerHand, card)
	if err != nil {
		return err
	}

	removeCard(playerHand, index)
	addCard(&CurrentGame.Board, card)

	return nil
}

func AddPlayer(name string) (*PlayerState, error) {
	if len(CurrentGame.Deck) < NumCardsInHand {
		return nil, errors.New("deck not big enough to make a new hand")
	}

	if len(CurrentGame.Players) == MaxPlayers {
		return nil, errors.New("no more new players can be added")
	}

	newPlayer, err := CreatePlayer(name)

	if err != nil {
		return nil, errors.New("error creating player")
	}

	CurrentGame = Game{
		Board:   CurrentGame.Board,
		Players: append(CurrentGame.Players, *newPlayer),
		Deck:    CurrentGame.Deck,
	}

	return newPlayer, nil
}
