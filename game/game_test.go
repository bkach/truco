package game

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_CreateGame_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()

	assert.NoError(t, err)
	// Check that the game state is as expected
	game := Game{
		ID:      "game_0",
		Board:   []Card{},
		Players: []PlayerState{},
		deck: []Card{
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

	t.Cleanup(func() {
		var games []Game
		Games = games
	})

	assert.Equal(t, []Game{game}, Games)
}

func Test_CreateMultipleGames_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	game0 := Game{
		ID:      "game_0",
		Board:   []Card{},
		Players: []PlayerState{},
		deck: []Card{
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

	game1 := Game{
		ID:      "game_1",
		Board:   []Card{},
		Players: []PlayerState{},
		deck: []Card{
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

	t.Cleanup(func() {
		var games []Game
		Games = games
	})

	assert.Equal(t, []Game{game0, game1}, Games)
}

func Test_CreateMultipleGamesAndAddUser_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	err = AddPlayer(Games, "game_1", "boris")
	assert.NoError(t, err)

	game0 := Game{
		ID:      "game_0",
		Board:   []Card{},
		Players: []PlayerState{},
		deck: []Card{
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

	game1 := Game{
		ID:    "game_1",
		Board: []Card{},
		Players: []PlayerState{
			{
				Info: PlayerInfo{
					Name: "boris",
					ID:   "player_boris",
				},
				Hand: []Card{
					{Value: 1, House: "gold"},
					{Value: 1, House: "cups"},
					{Value: 1, House: "spades"},
				},
			},
		},
		deck: []Card{
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

	t.Cleanup(func() {
		var games []Game
		Games = games
	})

	assert.Equal(t, []Game{game0, game1}, Games)
}

func Test_CreateMultipleGamesAndGetGameIds_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	t.Cleanup(func() {
		var games []Game
		Games = games
	})

	assert.Equal(t, []string{"game_0", "game_1"}, GetGameIds(Games))
}

func Test_PlayCard_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	err = AddPlayer(Games, "game_0", "boris")
	assert.NoError(t, err)

	err = PlayCard(Games, Card{Value: 1, House: "gold"}, "game_0", "player_boris")
	assert.NoError(t, err)

	games := []Game{
		{
			ID: "game_0",
			Board: []Card{
				{Value: 1, House: "gold"},
			},
			Players: []PlayerState{
				{
					Info: PlayerInfo{
						Name: "boris",
						ID:   "player_boris",
					},
					Hand: []Card{
						{Value: 1, House: "cups"},
						{Value: 1, House: "spades"},
					},
				},
			},
			deck: []Card{
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
		},
	}

	t.Cleanup(func() {
		var games []Game
		Games = games
	})

	assert.Equal(t, games, Games)
}

//noinspection ALL
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}
