package game

const MAX_PLAYERS = 4
const NUM_CARDS_IN_HAND = 3

type Game struct {
	Board   []Card
	Players []PlayerState
	Deck    []Card
}

var CurrentGame Game

// For testing
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

	findCardErr := findAndRemoveCard(&CurrentGame.Players[playerIndex].Hand, card)

	if findCardErr != nil {
		return findCardErr
	}

	addCard(&CurrentGame.Board, card)

	return err
}
