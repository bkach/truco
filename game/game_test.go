package game

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NewGame(t *testing.T) {
	debugOn = false

	// Create a new game
	StartGame()

	// Check that the game state is as expected
	expectedGameState := Game{
		Board:        []Card{},
		PlayerStates: []PlayerState{},
	}

	assert.Equal(t, CurrentGame, expectedGameState)
}

func Test_AddPlayer(t *testing.T) {
	debugOn = false

	StartGame()

	_, _, err := AddPlayer("boris")

	assert.NoError(t, err)

	expectedGameState := Game{
		Board: []Card{},
		PlayerStates: []PlayerState{
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
	}

	assert.Equal(t, CurrentGame, expectedGameState)
}
