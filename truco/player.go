package truco

import (
	"errors"
	"github.com/gofrs/uuid"
)

type Player struct {
	Id                string `json:"id"`
	Name              string `json:"name"`
	Hand              []Card `json:"cards"`
	CardIndicesPlayed []int  `json:"cardIndicesPlayed"`
}

func createPlayer(name string) (*Player, error) {
	playerUUID, err := uuid.NewV4()

	if err != nil {
		return nil, errors.New("error making UUID")
	}

	// For testing
	var playerId string
	if debugOn {
		playerId = "player_" + name
	} else {
		playerId = "player_" + playerUUID.String()
	}

	return &Player{
		Id:   playerId,
		Name: name,
	}, nil
}

func dealPlayerIn(deck []Card, player *Player) (*Player, []Card) {
	var hand []Card

	var card Card
	var updatedDeck = deck
	for i := 0; i < NumCardsInHand; i++ {
		updatedDeck, card = popRandomCard(updatedDeck)
		hand = append(hand, card)
	}

	player.Hand = hand

	return player, updatedDeck
}
