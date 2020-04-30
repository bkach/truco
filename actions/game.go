package actions

type Game struct {
	Board        Cards
	PlayerStates PlayerStates
}

var currentDeck Cards
var currentGame Game

// For testing
var debugOn = false

func startGame() Game {
	currentDeck = buildDeck()
	currentGame = Game{
		Board:        Cards([]Card{}),
		PlayerStates: PlayerStates([]PlayerState{}),
	}
	return currentGame
}

func playCard(card Card, id string) (PlayerState, Game, error) {
	_, playerState, error := currentGame.PlayerStates.findPlayer(id)

	if error != nil {
		return PlayerState{}, currentGame, error
	}

	playerState.Hand.findAndRemoveCard(card)
	currentGame.Board.addCard(card)

	return *playerState, currentGame, error
}
