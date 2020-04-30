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

type Cards []Card

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

func (cards *Cards) findCard(card Card) (int, *Card) {
	for i, v := range *cards {
		if v.Value == card.Value && v.House == card.House {
			return i, &v
		}
	}
	return -1, nil
}

func (cards *Cards) removeCard(index int) {
	*cards = append((*cards)[:index], (*cards)[index+1:]...)
}

func (cards *Cards) addCard(card Card) {
	*cards = append(*cards, card)
}

func (cards *Cards) findAndRemoveCard(card Card) {
	fmt.Printf("\n\ncards %d\n\n", len(*cards))
	index, _ := cards.findCard(card)
	if index != -1 {
		cards.removeCard(index)
	}
}

func (cards *Cards) getAndRemoveRandomCard() Card {
	var r int
	highestIndex := len(*cards) - 1

	if highestIndex == 0 {
		*cards = Cards{}
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

	cards.removeCard(r)

	return card
}
