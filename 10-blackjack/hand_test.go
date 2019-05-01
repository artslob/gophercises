package main

import (
	"github.com/artslob/gophercises/09-deck/deck"
	"testing"
)

// TODO: move hand to separate package
// TODO: make score as field of hand, score should be recalculated on creation of hand and on each draw

func TestHandDraw(t *testing.T) {
	h := hand{}
	h.draw(deck.Card{Suit: deck.Club, Rank: deck.Ace})
	h.draw(deck.Card{Suit: deck.Heart, Rank: deck.Seven})
	if len(h) != 2 {
		t.Fatal("expected size of hand to be 2")
	}
	h.draw(deck.Card{Suit: deck.Diamond, Rank: deck.Jack})
	if len(h) != 3 {
		t.Fatal("expected size of hand to be 3")
	}
	last := h[2]
	if last.Suit != deck.Diamond || last.Rank != deck.Jack {
		t.Fatal("last card is wrong:", last)
	}
}

func TestHandScore(t *testing.T) {
	tables := []struct {
		hand         hand
		normal, soft int
	}{
		{
			hand: hand{
				deck.Card{Suit: deck.Club, Rank: deck.Ace},
				deck.Card{Suit: deck.Heart, Rank: deck.Seven},
			},
			normal: 18,
			soft:   8,
		},
		{
			hand: hand{
				deck.Card{Suit: deck.Club, Rank: deck.Ace},
				deck.Card{Suit: deck.Heart, Rank: deck.Ace},
				deck.Card{Suit: deck.Diamond, Rank: deck.Ace},
			},
			normal: 33,
			soft:   3,
		},
		{
			hand: hand{
				deck.Card{Suit: deck.Club, Rank: deck.Ace},
			},
			normal: 11,
			soft:   1,
		},
		{
			hand: hand{
				deck.Card{Suit: deck.Club, Rank: deck.Jack},
				deck.Card{Suit: deck.Club, Rank: deck.Queen},
				deck.Card{Suit: deck.Spade, Rank: deck.King},
			},
			normal: 30,
			soft:   30,
		},
		{
			hand: hand{
				deck.Card{Suit: deck.Club, Rank: deck.Two},
				deck.Card{Suit: deck.Club, Rank: deck.Three},
				deck.Card{Suit: deck.Club, Rank: deck.Four},
				deck.Card{Suit: deck.Club, Rank: deck.Five},
			},
			normal: 14,
			soft:   14,
		},
		{
			hand: hand{
				deck.Card{Suit: deck.Club, Rank: deck.Six},
				deck.Card{Suit: deck.Club, Rank: deck.Seven},
				deck.Card{Suit: deck.Club, Rank: deck.Eight},
				deck.Card{Suit: deck.Club, Rank: deck.Nine},
				deck.Card{Suit: deck.Club, Rank: deck.Ten},
			},
			normal: 40,
			soft:   40,
		},
	}
	for _, test := range tables {
		normal, soft := test.hand.getScores()
		if normal != test.normal || soft != test.soft {
			t.Fatalf("expected score: %d (%d), got: %d (%d)", test.normal, test.soft, normal, soft)
		}
	}
}
