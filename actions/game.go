package actions

type Game struct {
	Board        []Card
	PlayerStates []PlayerState
}

// TODO: Put in game
var currentDeck []Card
var currentGame Game

// For testing
var debugOn = false

func startGame() Game {
	currentDeck = buildDeck()
	currentGame = Game{
		Board:        []Card{},
		PlayerStates: []PlayerState{},
	}
	return currentGame
}

func playCard(card Card, id string) (PlayerState, Game, error) {
	_, playerState, error := findPlayer(&currentGame.PlayerStates, id)

	if error != nil {
		return PlayerState{}, currentGame, error
	}

	findAndRemoveCard(&playerState.Hand, card)
	addCard(&currentGame.Board, card)

	return *playerState, currentGame, error
}
