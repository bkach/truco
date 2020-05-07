package truco

import (
	"errors"
	"sort"
)

const MaxPlayers = 4
const NumCardsInHand = 3

// This is the only file that can manipulate the global state

func CreateGameAndAddToGames() (string, error) {
	newGame, err := createGame(Games)

	if err != nil {
		return "", err
	}

	Games = append(Games, *newGame)

	return newGame.Id, nil
}

func FindGameWithId(id string) (int, *Game, error) {
	index := sort.Search(len(Games), func(i int) bool {
		return Games[i].Id == id
	})

	if index == len(Games) {
		return 0, nil, errors.New("game with id " + id + " not found")
	}

	return index, &Games[index], nil
}

func AddPlayer(gameId string, name string) (string, error) {
	gameIndex, game, err := FindGameWithId(gameId)

	if err != nil {
		return "", err
	}

	if len(game.Deck) < NumCardsInHand {
		return "", errors.New("deck not big enough to make a new hand")
	}

	if len(game.Players) == MaxPlayers {
		return "", errors.New("no more new players can be added")
	}

	newPlayerId, newPlayerState, err := createPlayer(name)

	if err != nil {
		return "", err
	}

	game.Players[newPlayerId] = *newPlayerState

	Games[gameIndex] = *game

	return newPlayerId, nil
}

func PlayCard(gameId string, playerId string, card Card) error {
	gameIndex, game, err := FindGameWithId(gameId)

	if err != nil {
		return err
	}

	playerHand := game.Players[playerId].Hand
	board := game.Board

	index, err := findCardIndex(playerHand, card)
	if err != nil {
		return err
	}

	playerHand = removeCard(playerHand, index)
	board = addCard(board, card)

	game.Players[playerId] = PlayerState{
		Name: game.Players[playerId].Name,
		Hand: playerHand, // Updated value
	}

	newGame := Game{
		Id:      gameId,
		Board:   board, // Updated value
		Players: game.Players,
		Deck:    game.Deck,
	}

	Games[gameIndex] = newGame

	return nil
}

func DealCards(gameId string) error {
	gameIndex, game, err := FindGameWithId(gameId)

	if err != nil {
		return err
	}

	if len(game.Deck) < NumCardsInHand*len(game.Players) {
		return errors.New("deck not big enough to deal cards")
	}

	newDeck := game.Deck
	for playerId, player := range game.Players {
		player, deck := dealPlayerIn(newDeck, &player)
		newDeck = deck

		game.Players[playerId] = *player
	}

	updatedGame := Game{
		Id:      game.Id,
		Board:   game.Board,
		Players: game.Players,
		Deck:    newDeck,
	}

	Games[gameIndex] = updatedGame

	return nil
}
