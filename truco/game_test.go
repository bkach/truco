package truco

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_CreateGame_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	game := Game{
		Board:   []Card{},
		Id:      "game_0",
		Players: []Player{},
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

	assert.Equal(t, []Game{game}, Games)
}

func Test_CreateMultipleGames_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = CreateGameAndAddToGames()
	assert.NoError(t, err)

	game0 := Game{
		Id:      "game_0",
		Board:   []Card{},
		Players: []Player{},
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
		Id:      "game_1",
		Board:   []Card{},
		Players: []Player{},
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

	assert.Equal(t, []Game{game0, game1}, Games)
}

func Test_CreateMultipleGamesAndAddUser_HasExpectedState(t *testing.T) {
	debugOn = true

	_, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	gameId1, err := CreateGameAndAddToGames()
	assert.NoError(t, err)

	_, err = AddPlayer(gameId1, "boris")
	assert.NoError(t, err)

	err = DealCards(gameId1)
	assert.NoError(t, err)

	game0 := Game{
		Id:      "game_0",
		Board:   []Card{},
		Players: []Player{},
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
		Id:    "game_1",
		Board: []Card{},
		Players: []Player{
			{
				Id: "player_boris",
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

	assert.Equal(t, []Game{game0, game1}, Games)
}

//func Test_PlayCard_HasExpectedState(t *testing.T) {
//	debugOn = true
//
//	gameId, err := CreateGameAndAddToGames()
//	assert.NoError(t, err)
//
//	_, err = AddPlayer(gameId, "boris")
//	assert.NoError(t, err)
//
//	_, err = AddPlayer(gameId, "papi")
//	assert.NoError(t, err)
//
//	_, err = AddPlayer(gameId, "claudio")
//	assert.NoError(t, err)
//
//	_, err = AddPlayer(gameId, "jorge")
//	assert.NoError(t, err)
//
//	err = DealCards(gameId)
//	assert.NoError(t, err)
//
//	err = PlayCard(gameId, "player_boris", Card{Value: 1, House: "gold"})
//
//	game := Game{
//		Id: "game_0",
//		Board: []Card{
//			{Value: 1, House: "gold"},
//		},
//		Players: []Player{
//			{
//				Id: "player_boris",
//				Name: "boris",
//				Hand: []Card{
//					{Value: 1, House: "cups"},
//					{Value: 1, House: "spades"},
//				},
//			},
//			{
//				Id: "player_papi",
//				Name: "papi",
//				Hand: []Card{
//					{Value: 1, House: "clubs"},
//					{Value: 2, House: "gold"},
//					{Value: 2, House: "cups"},
//				},
//			},
//			{
//				Id: "player_claudio",
//				Name: "claudio",
//				Hand: []Card{
//					{Value: 2, House: "spades"},
//					{Value: 2, House: "clubs"},
//					{Value: 3, House: "gold"},
//				},
//			},
//			{
//				Id: "player_jorge",
//				Name: "jorge",
//				Hand: []Card{
//					{Value: 3, House: "cups"},
//					{Value: 3, House: "spades"},
//					{Value: 3, House: "clubs"},
//				},
//			},
//		},
//		Deck: []Card{
//			{Value: 4, House: "gold"},
//			{Value: 4, House: "cups"},
//			{Value: 4, House: "spades"},
//			{Value: 4, House: "clubs"},
//			{Value: 5, House: "gold"},
//			{Value: 5, House: "cups"},
//			{Value: 5, House: "spades"},
//			{Value: 5, House: "clubs"},
//			{Value: 6, House: "gold"},
//			{Value: 6, House: "cups"},
//			{Value: 6, House: "spades"},
//			{Value: 6, House: "clubs"},
//			{Value: 7, House: "gold"},
//			{Value: 7, House: "cups"},
//			{Value: 7, House: "spades"},
//			{Value: 7, House: "clubs"},
//			{Value: 10, House: "gold"},
//			{Value: 10, House: "cups"},
//			{Value: 10, House: "spades"},
//			{Value: 10, House: "clubs"},
//			{Value: 11, House: "gold"},
//			{Value: 11, House: "cups"},
//			{Value: 11, House: "spades"},
//			{Value: 11, House: "clubs"},
//			{Value: 12, House: "gold"},
//			{Value: 12, House: "cups"},
//			{Value: 12, House: "spades"},
//			{Value: 12, House: "clubs"},
//		},
//	}
//
//	cleanup(t)
//
//	assert.Equal(t, []Game{ game }, Games)
//}

func cleanup(t *testing.T) {
	t.Cleanup(func() {
		Games = []Game{}
	})
}
