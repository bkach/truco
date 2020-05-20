package truco

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateGame_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames("game0")
	assert.NoError(t, err)

	game := Game{
		Name:    "game0",
		Id:      "game_0",
		Players: []Player{},
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

	cleanup(t)

	assert.Equal(t, []Game{game}, Games)
}

func Test_CreateMultipleGames_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames("game0")
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames("game1")
	assert.NoError(t, err)

	game0 := Game{
		Name:    "game0",
		Id:      "game_0",
		Players: []Player{},
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
		Name:    "game1",
		Id:      "game_1",
		Players: []Player{},
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

	cleanup(t)

	assert.Equal(t, []Game{game0, game1}, Games)
}

func Test_CreateMultipleGamesAndAddUser_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames("game0")
	assert.NoError(t, err)

	game, err := CreateGameAndAddToGames("game1")
	assert.NoError(t, err)

	_, err = CreatePlayer(game.Id, "boris")
	assert.NoError(t, err)

	err = DealCards(game.Id)
	assert.NoError(t, err)

	game0 := Game{
		Name:    "game0",
		Id:      "game_0",
		Players: []Player{},
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
		Name: "game1",
		Id:   "game_1",
		Players: []Player{
			{
				Id:   "player_boris",
				Name: "boris",
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

	cleanup(t)

	assert.Equal(t, []Game{game0, game1}, Games)
}

func Test_PlayCards_HasExpectedState(t *testing.T) {
	debugOn = true

	game, err := CreateGameAndAddToGames("game0")
	assert.NoError(t, err)

	_, err = CreatePlayer(game.Id, "boris")
	assert.NoError(t, err)

	_, err = CreatePlayer(game.Id, "papi")
	assert.NoError(t, err)

	_, err = CreatePlayer(game.Id, "claudio")
	assert.NoError(t, err)

	_, err = CreatePlayer(game.Id, "jorge")
	assert.NoError(t, err)

	err = DealCards(game.Id)
	assert.NoError(t, err)

	err = PlayCard(game.Id, "player_boris", Card{Value: 1, House: "gold"})
	assert.NoError(t, err)

	err = PlayCard(game.Id, "player_boris", Card{Value: 1, House: "spades"})
	assert.NoError(t, err)

	err = PlayCard(game.Id, "player_boris", Card{Value: 1, House: "cups"})
	assert.NoError(t, err)

	err = PlayCard(game.Id, "player_papi", Card{Value: 2, House: "gold"})
	assert.NoError(t, err)

	game1 := Game{
		Name: "game0",
		Id:   "game_0",
		Players: []Player{
			{
				Id:   "player_boris",
				Name: "boris",
				Hand: []Card{
					{Value: 1, House: "gold"},
					{Value: 1, House: "cups"},
					{Value: 1, House: "spades"},
				},
				CardIndicesPlayed: []int{
					0, 2, 1,
				},
			},
			{
				Id:   "player_papi",
				Name: "papi",
				Hand: []Card{
					{Value: 1, House: "clubs"},
					{Value: 2, House: "gold"},
					{Value: 2, House: "cups"},
				},
				CardIndicesPlayed: []int{
					1,
				},
			},
			{
				Id:   "player_claudio",
				Name: "claudio",
				Hand: []Card{
					{Value: 2, House: "spades"},
					{Value: 2, House: "clubs"},
					{Value: 3, House: "gold"},
				},
			},
			{
				Id:   "player_jorge",
				Name: "jorge",
				Hand: []Card{
					{Value: 3, House: "cups"},
					{Value: 3, House: "spades"},
					{Value: 3, House: "clubs"},
				},
			},
		},
		deck: []Card{
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

	assert.Equal(t, []Game{game1}, Games)
}

func Test_PlayCard_SameTime_HasError(t *testing.T) {
	debugOn = true

	game, err := CreateGameAndAddToGames("game0")
	assert.NoError(t, err)

	_, err = CreatePlayer(game.Id, "boris")
	assert.NoError(t, err)

	err = DealCards(game.Id)
	assert.NoError(t, err)

	err = PlayCard(game.Id, "player_boris", Card{Value: 1, House: "gold"})
	assert.NoError(t, err)

	// Play same card
	err = PlayCard(game.Id, "player_boris", Card{Value: 1, House: "gold"})
	if assert.Error(t, err) {
		assert.Equal(t,
			errors.New("card {Value:1 House:gold} already played in this hand, cannot play card again"),
			err,
		)
	}
}

// TODO: This could be split into much smaller tests
func Test_GameChangeListener_WorksAsExpected(t *testing.T) {
	debugOn = true

	id1 := "gameChangeListenerId1"
	listener1Counter := 0

	id2 := "gameChangeListenerId2"
	listener2Counter := 0

	RegisterGameChangeListener(id1, func() {
		listener1Counter++
	})

	RegisterGameChangeListener(id2, func() {
		listener2Counter++
	})

	game, err := CreateGameAndAddToGames("game0")
	assert.NoError(t, err)

	assert.Equal(t, 1, listener1Counter)
	assert.Equal(t, 1, listener2Counter)

	player, err := CreatePlayer(game.Id, "boris")
	assert.NoError(t, err)

	assert.Equal(t, 2, listener1Counter)
	assert.Equal(t, 2, listener2Counter)

	err = DealCards(game.Id)
	assert.NoError(t, err)

	assert.Equal(t, 3, listener1Counter)
	assert.Equal(t, 3, listener2Counter)

	err = PlayCard(game.Id, "player_boris", Card{Value: 1, House: "gold"})
	assert.NoError(t, err)

	assert.Equal(t, 4, listener1Counter)
	assert.Equal(t, 4, listener2Counter)

	err = DeletePlayer(game.Id, player.Id)
	assert.NoError(t, err)

	assert.Equal(t, 5, listener1Counter)
	assert.Equal(t, 5, listener2Counter)

	err = DeleteGame(game.Id)
	assert.NoError(t, err)

	assert.Equal(t, 6, listener1Counter)
	assert.Equal(t, 6, listener2Counter)

	RemoveGameChangeListener(id2)
	assert.Equal(t, 1, len(gameChangeListeners))

	_, err = CreateGameAndAddToGames("game0")
	assert.NoError(t, err)

	assert.Equal(t, 7, listener1Counter)
	assert.Equal(t, 6, listener2Counter)
}

func Test_Envido(t *testing.T) {
	e1 := calculateEnvido([]Card{
		{
			Value: 1,
			House: Gold,
		},
		{
			Value: 2,
			House: Gold,
		},
		{
			Value: 1,
			House: Spades,
		},
	})

	assert.Equal(t, 23, e1)

	e2 := calculateEnvido([]Card{
		{
			Value: 12,
			House: Gold,
		},
		{
			Value: 2,
			House: Gold,
		},
		{
			Value: 1,
			House: Spades,
		},
	})

	assert.Equal(t, 22, e2)

	e3 := calculateEnvido([]Card{
		{
			Value: 12,
			House: Gold,
		},
		{
			Value: 8,
			House: Spades,
		},
		{
			Value: 1,
			House: Clubs,
		},
	})

	assert.Equal(t, 8, e3)
}

func cleanup(t *testing.T) {
	t.Cleanup(func() {
		Games = []Game{}
	})
}
