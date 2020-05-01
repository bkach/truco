package game

import (
	"errors"
	"github.com/gofrs/uuid"
	"strconv"
)

const MaxPlayers = 4
const NumCardsInHand = 3

type Game struct {
	ID      string
	Board   []Card
	Players []PlayerState
	deck    []Card
}

// Global state
var Games []Game
var debugOn = false

func CreateGameAndAddToGames() (string, error) {
	newGame, err := createGame()

	if err != nil {
		return "", err
	}

	Games = append(Games, *newGame)

	return newGame.ID, nil
}

func createGame() (*Game, error) {
	gameUUID, err := uuid.NewV4()

	if err != nil {
		return nil, err
	}

	// For testing
	var gameId string
	if debugOn {
		gameId = "game_" + strconv.Itoa(len(Games))
	} else {
		gameId = "game_" + gameUUID.String()
	}

	game := Game{
		Board:   []Card{},
		Players: []PlayerState{},
		deck:    buildDeck(),
		ID:      gameId,
	}

	return &game, nil
}

func FindGameIndex(games []Game, gameId string) (int, error) {
	for index, game := range games {
		if game.ID == gameId {
			return index, nil
		}
	}
	return -1, errors.New("cannot find game " + gameId)
}

func GetGameIds(games []Game) []string {
	var gameIds []string
	for _, selectedGame := range games {
		gameIds = append(gameIds, selectedGame.ID)
	}

	return gameIds
}

func PlayCard(games []Game, card Card, gameId string, playerId string) error {
	gameIndex, err := FindGameIndex(games, gameId)

	if err != nil {
		return err
	}

	playerIndex, _, err := FindPlayer(games, gameId, playerId)

	if err != nil {
		return err
	}

	playerHand := &games[gameIndex].Players[playerIndex].Hand

	index, err := findCardIndex(playerHand, card)
	if err != nil {
		return err
	}

	removeCard(playerHand, index)
	addCard(&games[gameIndex].Board, card)

	return nil
}

func AddPlayer(games []Game, gameId string, name string) error {
	gameIndex, err := FindGameIndex(games, gameId)

	if err != nil {
		return err
	}

	if len(games[gameIndex].deck) < NumCardsInHand {
		return errors.New("deck not big enough to make a new hand")
	}

	if len(games[gameIndex].Players) == MaxPlayers {
		return errors.New("no more new players can be added")
	}

	newPlayer, err := createPlayer(&games[gameIndex], name)

	if err != nil {
		return errors.New("error creating player")
	}

	games[gameIndex].Players = append(games[gameIndex].Players, *newPlayer)

	return nil
}
