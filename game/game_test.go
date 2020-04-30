package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewGame_HasExpectedState(t *testing.T) {
	debugOn = true

	// Create a new game
	GlobalGame = StartGame()

	// Check that the game state is as expected
	expectedGameState := Game{
		Board:   []Card{},
		Players: []PlayerState{},
		Deck: []Card{
			{Value: 1, House: "gold"},
			{Value: 1, House: "cups"},
			{Value: 1, House: "spades"},
			{Value: 1, House: "clubs"},
			{Value: 2, House: "gold"},
			{Value: 2, House: "cups"},
			{Value: 2, House: "spades"},
			{Value: 2, House: "clubs"},
			{Value: 3, House: "gold"},
			{Value: 3, House: "cups"},
			{Value: 3, House: "spades"},
			{Value: 3, House: "clubs"},
			{Value: 4, House: "gold"},
			{Value: 4, House: "cups"},
			{Value: 4, House: "spades"},
			{Value: 4, House: "clubs"},
			{Value: 5, House: "gold"},
			{Value: 5, House: "cups"},
			{Value: 5, House: "spades"},
			{Value: 5, House: "clubs"},
			{Value: 6, House: "gold"},
			{Value: 6, House: "cups"},
			{Value: 6, House: "spades"},
			{Value: 6, House: "clubs"},
			{Value: 7, House: "gold"},
			{Value: 7, House: "cups"},
			{Value: 7, House: "spades"},
			{Value: 7, House: "clubs"},
			{Value: 10, House: "gold"},
			{Value: 10, House: "cups"},
			{Value: 10, House: "spades"},
			{Value: 10, House: "clubs"},
			{Value: 11, House: "gold"},
			{Value: 11, House: "cups"},
			{Value: 11, House: "spades"},
			{Value: 11, House: "clubs"},
			{Value: 12, House: "gold"},
			{Value: 12, House: "cups"},
			{Value: 12, House: "spades"},
			{Value: 12, House: "clubs"},
		},
	}

	assert.Equal(t, expectedGameState, GlobalGame)
}

func Test_AddPlayer_HasExpectedState(t *testing.T) {
	debugOn = true

	GlobalGame = StartGame()

	addedPlayer, err := AddPlayer(&GlobalGame, "boris")

	assert.NoError(t, err)

	expectedAddedPlayer := PlayerState{
		Info: PlayerInfo{
			Name: "boris",
			ID:   "player_boris",
		},
		Hand: []Card{
			{
				Value: 1,
				House: Gold,
			},
			{
				Value: 1,
				House: Cups,
			},
			{
				Value: 1,
				House: Spades,
			},
		},
	}

	assert.Equal(t, &expectedAddedPlayer, addedPlayer)

	expectedGameState := Game{
		Board: []Card{},
		Players: []PlayerState{
			expectedAddedPlayer,
		},
		Deck: []Card{
			{Value: 1, House: "clubs"},
			{Value: 2, House: "gold"},
			{Value: 2, House: "cups"},
			{Value: 2, House: "spades"},
			{Value: 2, House: "clubs"},
			{Value: 3, House: "gold"},
			{Value: 3, House: "cups"},
			{Value: 3, House: "spades"},
			{Value: 3, House: "clubs"},
			{Value: 4, House: "gold"},
			{Value: 4, House: "cups"},
			{Value: 4, House: "spades"},
			{Value: 4, House: "clubs"},
			{Value: 5, House: "gold"},
			{Value: 5, House: "cups"},
			{Value: 5, House: "spades"},
			{Value: 5, House: "clubs"},
			{Value: 6, House: "gold"},
			{Value: 6, House: "cups"},
			{Value: 6, House: "spades"},
			{Value: 6, House: "clubs"},
			{Value: 7, House: "gold"},
			{Value: 7, House: "cups"},
			{Value: 7, House: "spades"},
			{Value: 7, House: "clubs"},
			{Value: 10, House: "gold"},
			{Value: 10, House: "cups"},
			{Value: 10, House: "spades"},
			{Value: 10, House: "clubs"},
			{Value: 11, House: "gold"},
			{Value: 11, House: "cups"},
			{Value: 11, House: "spades"},
			{Value: 11, House: "clubs"},
			{Value: 12, House: "gold"},
			{Value: 12, House: "cups"},
			{Value: 12, House: "spades"},
			{Value: 12, House: "clubs"},
		},
	}

	assert.Equal(t, expectedGameState, GlobalGame)
}

func Test_PlayCard_HasExpectedState(t *testing.T) {
	debugOn = true

	GlobalGame = StartGame()

	_, err := AddPlayer(&GlobalGame, "boris")

	assert.NoError(t, err)

	playCardErr := PlayCard(&GlobalGame, Card{1, Spades}, "player_boris")

	assert.NoError(t, playCardErr)

	expectedGameState := Game{
		Board: []Card{
			{
				Value: 1,
				House: Spades,
			},
		},
		Players: []PlayerState{
			{
				Info: PlayerInfo{
					Name: "boris",
					ID:   "player_boris",
				},
				Hand: []Card{
					{
						Value: 1,
						House: Gold,
					},
					{
						Value: 1,
						House: Cups,
					},
				},
			},
		},
		Deck: []Card{
			{Value: 1, House: "clubs"},
			{Value: 2, House: "gold"},
			{Value: 2, House: "cups"},
			{Value: 2, House: "spades"},
			{Value: 2, House: "clubs"},
			{Value: 3, House: "gold"},
			{Value: 3, House: "cups"},
			{Value: 3, House: "spades"},
			{Value: 3, House: "clubs"},
			{Value: 4, House: "gold"},
			{Value: 4, House: "cups"},
			{Value: 4, House: "spades"},
			{Value: 4, House: "clubs"},
			{Value: 5, House: "gold"},
			{Value: 5, House: "cups"},
			{Value: 5, House: "spades"},
			{Value: 5, House: "clubs"},
			{Value: 6, House: "gold"},
			{Value: 6, House: "cups"},
			{Value: 6, House: "spades"},
			{Value: 6, House: "clubs"},
			{Value: 7, House: "gold"},
			{Value: 7, House: "cups"},
			{Value: 7, House: "spades"},
			{Value: 7, House: "clubs"},
			{Value: 10, House: "gold"},
			{Value: 10, House: "cups"},
			{Value: 10, House: "spades"},
			{Value: 10, House: "clubs"},
			{Value: 11, House: "gold"},
			{Value: 11, House: "cups"},
			{Value: 11, House: "spades"},
			{Value: 11, House: "clubs"},
			{Value: 12, House: "gold"},
			{Value: 12, House: "cups"},
			{Value: 12, House: "spades"},
			{Value: 12, House: "clubs"},
		},
	}

	assert.Equal(t, expectedGameState, GlobalGame)
}

func Test_PlayInvalidCard_ThrowsErrorAndDoesNotChangeState(t *testing.T) {
	debugOn = true

	GlobalGame = StartGame()

	_, err := AddPlayer(&GlobalGame, "boris")

	assert.NoError(t, err)

	playCardErr := PlayCard(&GlobalGame, Card{10, Spades}, "player_boris")

	assert.Error(t, playCardErr)

	expectedGameState := Game{
		Board: []Card{},
		Players: []PlayerState{
			{
				Info: PlayerInfo{
					Name: "boris",
					ID:   "player_boris",
				},
				Hand: []Card{
					{
						Value: 1,
						House: Gold,
					},
					{
						Value: 1,
						House: Cups,
					},
					{
						Value: 1,
						House: Spades,
					},
				},
			},
		},
		Deck: []Card{
			{Value: 1, House: "clubs"},
			{Value: 2, House: "gold"},
			{Value: 2, House: "cups"},
			{Value: 2, House: "spades"},
			{Value: 2, House: "clubs"},
			{Value: 3, House: "gold"},
			{Value: 3, House: "cups"},
			{Value: 3, House: "spades"},
			{Value: 3, House: "clubs"},
			{Value: 4, House: "gold"},
			{Value: 4, House: "cups"},
			{Value: 4, House: "spades"},
			{Value: 4, House: "clubs"},
			{Value: 5, House: "gold"},
			{Value: 5, House: "cups"},
			{Value: 5, House: "spades"},
			{Value: 5, House: "clubs"},
			{Value: 6, House: "gold"},
			{Value: 6, House: "cups"},
			{Value: 6, House: "spades"},
			{Value: 6, House: "clubs"},
			{Value: 7, House: "gold"},
			{Value: 7, House: "cups"},
			{Value: 7, House: "spades"},
			{Value: 7, House: "clubs"},
			{Value: 10, House: "gold"},
			{Value: 10, House: "cups"},
			{Value: 10, House: "spades"},
			{Value: 10, House: "clubs"},
			{Value: 11, House: "gold"},
			{Value: 11, House: "cups"},
			{Value: 11, House: "spades"},
			{Value: 11, House: "clubs"},
			{Value: 12, House: "gold"},
			{Value: 12, House: "cups"},
			{Value: 12, House: "spades"},
			{Value: 12, House: "clubs"},
		},
	}

	assert.Equal(t, expectedGameState, GlobalGame)
}

func Test_PlayInvalidPlayer_ThrowsErrorAndDoesNotChangeState(t *testing.T) {
	debugOn = true

	GlobalGame = StartGame()

	_, err := AddPlayer(&GlobalGame, "boris")

	assert.NoError(t, err)

	playCardErr := PlayCard(&GlobalGame, Card{1, Spades}, "some other player")

	assert.Error(t, playCardErr)

	expectedGameState := Game{
		Board: []Card{},
		Players: []PlayerState{
			{
				Info: PlayerInfo{
					Name: "boris",
					ID:   "player_boris",
				},
				Hand: []Card{
					{
						Value: 1,
						House: Gold,
					},
					{
						Value: 1,
						House: Cups,
					},
					{
						Value: 1,
						House: Spades,
					},
				},
			},
		},
		Deck: []Card{
			{Value: 1, House: "clubs"},
			{Value: 2, House: "gold"},
			{Value: 2, House: "cups"},
			{Value: 2, House: "spades"},
			{Value: 2, House: "clubs"},
			{Value: 3, House: "gold"},
			{Value: 3, House: "cups"},
			{Value: 3, House: "spades"},
			{Value: 3, House: "clubs"},
			{Value: 4, House: "gold"},
			{Value: 4, House: "cups"},
			{Value: 4, House: "spades"},
			{Value: 4, House: "clubs"},
			{Value: 5, House: "gold"},
			{Value: 5, House: "cups"},
			{Value: 5, House: "spades"},
			{Value: 5, House: "clubs"},
			{Value: 6, House: "gold"},
			{Value: 6, House: "cups"},
			{Value: 6, House: "spades"},
			{Value: 6, House: "clubs"},
			{Value: 7, House: "gold"},
			{Value: 7, House: "cups"},
			{Value: 7, House: "spades"},
			{Value: 7, House: "clubs"},
			{Value: 10, House: "gold"},
			{Value: 10, House: "cups"},
			{Value: 10, House: "spades"},
			{Value: 10, House: "clubs"},
			{Value: 11, House: "gold"},
			{Value: 11, House: "cups"},
			{Value: 11, House: "spades"},
			{Value: 11, House: "clubs"},
			{Value: 12, House: "gold"},
			{Value: 12, House: "cups"},
			{Value: 12, House: "spades"},
			{Value: 12, House: "clubs"},
		},
	}

	assert.Equal(t, expectedGameState, GlobalGame)
}
