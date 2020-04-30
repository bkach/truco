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
var GlobalGame Game
var debugOn = false

func StartGame() Game {
	game := Game{
		Board:   []Card{},
		Players: []PlayerState{},
		Deck:    buildDeck(),
	}
	return game
}

func PlayCard(game *Game, card Card, playerId string) error {
	playerIndex, _, err := FindPlayer(game.Players, playerId)

	if err != nil {
		return err
	}

	playerHand := &game.Players[playerIndex].Hand

	index, err := findCardIndex(playerHand, card)
	if err != nil {
		return err
	}

	removeCard(playerHand, index)
	addCard(&game.Board, card)

	return nil
}

func AddPlayer(game *Game, name string) (*PlayerState, error) {
	if len(game.Deck) < NumCardsInHand {
		return nil, errors.New("deck not big enough to make a new hand")
	}

	if len(game.Players) == MaxPlayers {
		return nil, errors.New("no more new players can be added")
	}

	newPlayer, err := createPlayer(game, name)

	if err != nil {
		return nil, errors.New("error creating player")
	}

	game.Players = append(game.Players, *newPlayer)

	return newPlayer, nil
}
