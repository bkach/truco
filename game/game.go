package game

import (
	"errors"
	"github.com/gofrs/uuid"
	"strconv"
)

const MaxPlayers = 4
const NumCardsInHand = 3

type Game struct {
	Board   []Card
	Players map[string]PlayerState
	Deck    []Card
}

// Global state
var Games = map[string]Game{}
var debugOn = false

func CreateGameAndAddToGames() (string, error) {
	gameId, newGame, err := createGame()

	if err != nil {
		return "", err
	}

	Games[gameId] = *newGame

	return gameId, nil
}

func createGame() (string, *Game, error) {
	gameUUID, err := uuid.NewV4()

	if err != nil {
		return "", nil, err
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
		Players: map[string]PlayerState{},
		Deck:    buildDeck(),
	}

	return gameId, &game, nil
}

func GetGameIds(games map[string]Game) []string {
	gameIds := make([]string, 0, len(games))
	for gameId := range games {
		gameIds = append(gameIds, gameId)
	}

	return gameIds
}

func PlayCard(games map[string]Game, card Card, gameId string, playerId string) error {
	playerHand := games[gameId].Players[playerId].Hand
	board := games[gameId].Board

	index, err := findCardIndex(playerHand, card)
	if err != nil {
		return err
	}

	removeCard(&playerHand, index)
	addCard(&board, card)

	// Update player in the map
	games[gameId].Players[playerId] = PlayerState{
		Name: games[gameId].Players[playerId].Name,
		Hand: playerHand, // Updated value
	}

	// Update game
	games[gameId] = Game{
		Board:   board, // Updated value
		Players: games[gameId].Players,
		Deck:    games[gameId].Deck,
	}

	return nil
}

func AddPlayer(games map[string]Game, gameId string, name string) error {
	if len(games[gameId].Deck) < NumCardsInHand {
		return errors.New("deck not big enough to make a new hand")
	}

	if len(games[gameId].Players) == MaxPlayers {
		return errors.New("no more new players can be added")
	}

	newPlayerId, newPlayerState, err := createPlayer(name)

	if err != nil {
		return err
	}

	//newDeck := dealPlayerIn(games[gameId].Deck, newPlayerState)

	// Add player to the map
	games[gameId].Players[newPlayerId] = *newPlayerState

	// Update game
	//games[gameId] = Game{
	//	Board:   games[gameId].Board,
	//	Players: games[gameId].Players,
	//	Deck:    newDeck,
	//}

	return nil
}

func DealCards(games map[string]Game, gameId string) error {
	if len(games[gameId].Deck) < NumCardsInHand*len(games[gameId].Players) {
		return errors.New("deck not big enough to deal cards")
	}

	newDeck := games[gameId].Deck
	for playerId, player := range games[gameId].Players {
		player, deck := dealPlayerIn(newDeck, &player)
		newDeck = deck

		games[gameId].Players[playerId] = *player
	}

	games[gameId] = Game{
		Board:   games[gameId].Board,
		Players: games[gameId].Players,
		Deck:    newDeck,
	}

	return nil
}
