package game

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateGame_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	game := Game{
		Board:   []Card{},
		Players: map[string]PlayerState{},
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

	cleanup(t)

	assert.Equal(t, map[string]Game{
		"game_0": game,
	}, Games)
}

func Test_CreateMultipleGames_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	game0 := Game{
		Board:   []Card{},
		Players: map[string]PlayerState{},
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

	game1 := Game{
		Board:   []Card{},
		Players: map[string]PlayerState{},
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

	cleanup(t)

	assert.Equal(t, map[string]Game{"game_0": game0, "game_1": game1}, Games)
}

func Test_CreateMultipleGamesAndAddUser_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	err = AddPlayer(Games, "game_1", "boris")
	assert.NoError(t, err)

	err = DealCards(Games, "game_1")
	assert.NoError(t, err)

	game0 := Game{
		Board:   []Card{},
		Players: map[string]PlayerState{},
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

	game1 := Game{
		Board: []Card{},
		Players: map[string]PlayerState{
			"player_boris": {
				Name: "boris",
				Hand: []Card{
					{Value: 1, House: "gold"},
					{Value: 1, House: "cups"},
					{Value: 1, House: "spades"},
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

	cleanup(t)

	assert.Equal(t, map[string]Game{
		"game_0": game0,
		"game_1": game1,
	}, Games)
}

func Test_CreateMultipleGamesAndGetGameIds_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	cleanup(t)

	assert.Equal(t, []string{"game_0", "game_1", "game_2"}, GetGameIds(Games))
}

func Test_PlayCard_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	err = AddPlayer(Games, "game_0", "boris")
	assert.NoError(t, err)

	err = AddPlayer(Games, "game_0", "papi")
	assert.NoError(t, err)

	err = AddPlayer(Games, "game_0", "claudio")
	assert.NoError(t, err)

	err = AddPlayer(Games, "game_0", "jorge")
	assert.NoError(t, err)

	err = DealCards(Games, "game_0")
	assert.NoError(t, err)

	err = PlayCard(Games, Card{Value: 1, House: "gold"}, "game_0", "player_boris")

	game := Game{
		Board: []Card{
			{Value: 1, House: "gold"},
		},
		Players: map[string]PlayerState{
			"player_boris": {
				Name: "boris",
				Hand: []Card{
					{Value: 1, House: "cups"},
					{Value: 1, House: "spades"},
				},
			},
			"player_papi": {
				Name: "papi",
				Hand: []Card{
					{Value: 1, House: "clubs"},
					{Value: 2, House: "gold"},
					{Value: 2, House: "cups"},
				},
			},
			"player_claudio": {
				Name: "claudio",
				Hand: []Card{
					{Value: 2, House: "spades"},
					{Value: 2, House: "clubs"},
					{Value: 3, House: "gold"},
				},
			},
			"player_jorge": {
				Name: "jorge",
				Hand: []Card{
					{Value: 3, House: "cups"},
					{Value: 3, House: "spades"},
					{Value: 3, House: "clubs"},
				},
			},
		},
		Deck: []Card{
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

	cleanup(t)

	assert.Equal(t, map[string]Game{"game_0": game}, Games)
}

//noinspection ALL
func prettyPrint(i interface{}) string {
	s, _ := json.MarshalIndent(i, "", "\t")
	return string(s)
}

func cleanup(t *testing.T) {
	t.Cleanup(func() {
		Games = map[string]Game{}
	})
}
