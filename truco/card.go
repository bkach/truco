package truco

import (
	"errors"
	"fmt"
	"math/rand"
)

type House string

// The five houses of the spanish card deck
const (
	Gold   House = "gold"
	Cups   House = "cups"
	Spades House = "spades"
	Clubs  House = "clubs"
)

// A single playing card, values 1-12
type Card struct {
	Value int   `json:"value"`
	House House `json:"house"`
}

func buildDeck() []Card {
	var cards []Card
	for i := 1; i <= 12; i++ {
		// Truco doesn't use the 8 or 9 cards
		if i != 8 && i != 9 {
			cards = append(cards, Card{Value: i, House: Gold})
			cards = append(cards, Card{Value: i, House: Cups})
			cards = append(cards, Card{Value: i, House: Spades})
			cards = append(cards, Card{Value: i, House: Clubs})
		}
	}
	return cards
}

func findCardIndex(cards []Card, card Card) (int, error) {
	for i, v := range cards {
		if v.Value == card.Value && v.House == card.House {
			return i, nil
		}
	}

	msg := fmt.Sprintf("\nError: Cannot find card %+v", card)
	return -1, errors.New(msg)
}

func removeCard(c []Card, i int) []Card {
	return append((c)[:i], (c)[i+1:]...)
}

func popRandomCard(cards []Card) ([]Card, Card) {
	var randomIndex int
	var poppedCard Card
	highestIndex := len(cards) - 1

	var updatedCards []Card
	if highestIndex == 0 {
		poppedCard = (cards)[0]
		updatedCards = []Card{}
		return updatedCards, poppedCard
	} else {
		// For testing
		if debugOn {
			randomIndex = 0
		} else {
			randomIndex = rand.Intn(highestIndex)
		}
	}

	poppedCard = cards[randomIndex]

	updatedCards = removeCard(cards, randomIndex)

	return updatedCards, poppedCard
}
