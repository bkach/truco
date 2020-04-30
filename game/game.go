package game

type Game struct {
	Board        []Card
	PlayerStates []PlayerState
}

// TODO: Put in game
var currentDeck []Card
var CurrentGame Game

// For testing
var debugOn = false

func StartGame() Game {
	currentDeck = buildDeck()
	CurrentGame = Game{
		Board:        []Card{},
		PlayerStates: []PlayerState{},
	}
	return CurrentGame
}

func PlayCard(card Card, id string) (PlayerState, Game, error) {
	_, playerState, err := findPlayer(&CurrentGame.PlayerStates, id)

	if err != nil {
		return PlayerState{}, CurrentGame, err
	}

	findAndRemoveCard(&playerState.Hand, card)
	addCard(&CurrentGame.Board, card)

	return *playerState, CurrentGame, err
}
