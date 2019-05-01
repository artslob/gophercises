package main

import (
	"github.com/artslob/gophercises/09-deck/deck"
	"log"
	"testing"
)

func TestCardStringer(t *testing.T) {
	tables := []struct {
		card     deck.Card
		expected string
	}{
		{
			deck.Card{Suit: deck.Spade, Rank: deck.Ace},
			"Spade: Ace",
		},
		{
			deck.Card{Suit: deck.Diamond, Rank: deck.Seven},
			"Diamond: Seven",
		},
		{
			deck.Card{Suit: deck.Club, Rank: deck.King},
			"Club: King",
		},
	}
	for _, test := range tables {
		s := test.card.String()
		if s != test.expected {
			log.Fatalf("expcted string of '%s' card to be '%s'", s, test.expected)
		}
	}
}
