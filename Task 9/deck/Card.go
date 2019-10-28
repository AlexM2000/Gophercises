package deck

import (
	"math/rand"
	"sort"
)

type Suit uint8

type Rank uint8

type Card struct {
	Suit
	Rank
}

const DECKSIZE = 36

const (
	Spade = iota
	Diamond
	Club
	Heart
)

const (
	_ Rank = iota + 4
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
	Ace
)

const (
	MinRank = Six
	MaxRank = Ace
)

var suits = []Suit{Spade, Diamond, Club, Heart}

func (card Card) String() string {

}

func New() []Card {
	newDeck := make([]Card, DECKSIZE, DECKSIZE)
	for _, suit := range suits {
		for value := MinRank; value <= MaxRank; value++ {
			newDeck = append(newDeck, Card{Suit: suit, Rank: value})
		}
	}
	return newDeck
}

func DefaultSort(deck []Card) []Card {
	sort.Slice(deck, Less(deck))
	return deck
}

func CustomSort(sorter func(deck []Card) func(i, j int) bool) func(deck []Card) []Card {
	return func(deck []Card) []Card {
		sort.Slice(deck, sorter(deck))
		return deck
	}
}

func Less(deck []Card) func(i, j int) bool {
	return func(i, j int) bool {
		return deck[i].Rank < deck[j].Rank
	}
}

func Shuffle(deck []Card) []Card {
	rand.Shuffle(len(deck),
		func(i, j int) {
			deck[i], deck[j] = deck[j], deck[i]
		})
	return deck
}

func FilterOut(f func(card Card) bool) func(deck []Card) []Card {
	return func(deck []Card) []Card {
		var filteredDeck []Card
		for _, card := range deck {
			if !f(card) {
				filteredDeck = append(filteredDeck, card)
			}
		}
		return filteredDeck
	}
}
