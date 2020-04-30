package game

import (
	"errors"
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

func findCardIndex(cards *[]Card, card Card) (int, error) {
	for i, v := range *cards {
		if v.Value == card.Value && v.House == card.House {
			return i, nil
		}
	}
	return -1, errors.New("cannot find card")
}

func removeCard(cards *[]Card, index int) {
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

func addCard(cards *[]Card, card Card) {
	*cards = append(*cards, card)
}

func popRandomCard(cards *[]Card) Card {
	var r int
	highestIndex := len(*cards) - 1

	if highestIndex == 0 {
		*cards = []Card{}
		return (*cards)[0]
	} else {
		// For testing
		if debugOn {
			r = 0
		} else {
			r = rand.Intn(highestIndex)
		}
	}

	card := (*cards)[r]

	removeCard(cards, r)

	return card
}
