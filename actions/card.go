package actions

import (
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

func findCard(cards *[]Card, card Card) (int, *Card) {
	for i, v := range *cards {
		if v.Value == card.Value && v.House == card.House {
			return i, &v
		}
	}
	return -1, nil
}

func removeCard(cards *[]Card, index int) {
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

func addCard(cards *[]Card, card Card) {
	*cards = append(*cards, card)
}

func findAndRemoveCard(cards *[]Card, card Card) {
	fmt.Printf("\n\ncards %d\n\n", len(*cards))
	index, _ := findCard(cards, card)
	if index != -1 {
		removeCard(cards, index)
	}
}

func getAndRemoveRandomCard(cards *[]Card) Card {
	var r int
	highestIndex := len(*cards) - 1

	if highestIndex == 0 {
		*cards = []Card{}
		return (*cards)[0]
	} else {
		// For testing
		if debugOn {
			r = rand.Intn(highestIndex)
		} else {
			r = 0
		}
	}

	card := (*cards)[r]

	removeCard(cards, r)

	return card
}
