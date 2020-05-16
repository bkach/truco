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

// Calculates Envido
//
// How this is calculated (From https://en.wikipedia.org/wiki/Truco):
// - The score of a pair of the same suit is the sum of the values of the cards + 20, but considering that the King
//  (12s), the Knights (11s) and Sotas (10s) are worth 0.
// - If the player has no suit pair, then his puntos de envido is the value of his highest card, with Kings, Knights and
//  Sotas worth 0.
// - If playing without Flor, in case of having three cards of the same suit, the puntos de envido are those of the
//  highest pair of the hand.
func calculateEnvido(hand []Card) int {
	// We assume there are only three cards

	cardsByHouse := make(map[House][]Card)

	for _, v := range hand {
		cards := cardsByHouse[v.House]

		if cards == nil {
			cards = []Card{}
		}

		cards = append(cards, v)

		cardsByHouse[v.House] = cards
	}

	maxCardValue := 0

	for _, cardsInHouse := range cardsByHouse {
		// If two or more cards share the same house, add 20
		if len(cardsInHouse) > 1 {
			finalValue := 20
			for _, card := range cardsInHouse {
				// 10, 11, and 12s are worth 0
				if card.Value < 10 {
					finalValue += card.Value
				}
			}
			return finalValue
		} else { // else return highest card
			for _, card := range cardsInHouse {
				if card.Value < 10 && card.Value > maxCardValue {
					maxCardValue = card.Value
				}
			}
		}
	}

	return maxCardValue
}
