package truco

import (
	"errors"
)

const MaxPlayers = 4
const NumCardsInHand = 3

// This file contains all the actions a user can perform on a game. It is the only file that can manipulate the global
// state

func CreateGameAndAddToGames() (string, error) {
	newGame, err := createGame(Games)

	if err != nil {
		return "", err
	}

	Games = append(Games, *newGame)

	return newGame.Id, nil
}

func DeleteGame(gameId string) error {
	index, _, err := FindGameWithId(gameId)

	if err != nil {
		return err
	}

	Games = append(Games[:index], Games[index+1:]...)

	return nil
}

func NumGames() int {
	return len(Games)
}

func FindGameWithId(id string) (int, *Game, error) {
	gameIndex := -1
	for index, game := range Games {
		if game.Id == id {
			gameIndex = index
		}
	}

	if gameIndex == -1 {
		return 0, nil, errors.New("game with id \"" + id + "\" not found")
	}

	return gameIndex, &Games[gameIndex], nil
}

func FindPlayerWithId(game Game, id string) (int, *Player, error) {
	playerIndex := -1
	for index, player := range game.Players {
		if player.Id == id {
			playerIndex = index
		}
	}

	if playerIndex == -1 {
		return 0, nil, errors.New("player with id \"" + id + "\" not found in game " + game.Id)
	}

	return playerIndex, &game.Players[playerIndex], nil
}

func CreatePlayer(gameId string, name string) (string, error) {
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

	newPlayer, err := createPlayer(name)

	if err != nil {
		return "", err
	}

	game.Players = append(game.Players, *newPlayer)

	Games[gameIndex] = *game

	return newPlayer.Id, nil
}

func DeletePlayer(gameId string, playerId string) error {
	gameIndex, _, err := FindGameWithId(gameId)

	if err != nil {
		return err
	}

	playerIndex, _, err := FindPlayerWithId(Games[gameIndex], playerId)

	if err != nil {
		return err
	}

	players := Games[gameIndex].Players

	Games[gameIndex].Players = append(players[:playerIndex], players[playerIndex+1:]...)

	return nil
}

func PlayCard(gameId string, playerId string, card Card) error {
	gameIndex, game, err := FindGameWithId(gameId)

	if err != nil {
		return err
	}

	playerIndex, player, err := FindPlayerWithId(*game, playerId)

	if err != nil {
		return err
	}

	playerHand := player.Hand
	board := game.Board

	index, err := findCardIndex(playerHand, card)
	if err != nil {
		return err
	}

	playerHand = removeCard(playerHand, index)
	board = addCard(board, card)

	game.Players[playerIndex] = Player{
		Id:   game.Players[playerIndex].Id,
		Name: game.Players[playerIndex].Name,
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
