package main

import (
	"github.com/artslob/gophercises/09-deck/deck"
	"github.com/artslob/gophercises/10-blackjack/hand"
	"testing"
)

func TestHandDraw(t *testing.T) {
	h := hand.Hand{}
	if h.Size() != 0 {
		t.Fatal("expected size of hand to be 0")
	}
	h.Draw(deck.Card{Suit: deck.Club, Rank: deck.Ace})
	h.Draw(deck.Card{Suit: deck.Heart, Rank: deck.Seven})
	if h.Size() != 2 {
		t.Fatal("expected size of hand to be 2")
	}
	top := deck.Card{Suit: deck.Diamond, Rank: deck.Jack}
	h.Draw(top)
	if h.Size() != 3 {
		t.Fatal("expected size of hand to be 3")
	}
	if h.TopCard() != top {
		t.Fatalf("last card is wrong, expected: %s, got: %s", top, h.TopCard())
	}
}

func TestHandScore(t *testing.T) {
	tables := []struct {
		hand         hand.Hand
		normal, soft int
	}{
		{
			hand:   hand.Hand{Cards: nil},
			normal: 0,
			soft:   0,
		},
		{
			hand:   hand.New(),
			normal: 0,
			soft:   0,
		},
		{
			hand: hand.New(
				deck.Card{Suit: deck.Club, Rank: deck.Ace},
				deck.Card{Suit: deck.Heart, Rank: deck.Seven},
			),
			normal: 18,
			soft:   8,
		},
		{
			hand: hand.Hand{
				Cards: &[]deck.Card{
					{Suit: deck.Club, Rank: deck.Ace},
					{Suit: deck.Heart, Rank: deck.Seven},
				},
			},
			normal: 18,
			soft:   8,
		},
		{
			hand: hand.New(
				deck.Card{Suit: deck.Club, Rank: deck.Ace},
				deck.Card{Suit: deck.Heart, Rank: deck.Ace},
				deck.Card{Suit: deck.Diamond, Rank: deck.Ace},
			),
			normal: 33,
			soft:   3,
		},
		{
			hand: hand.New(
				deck.Card{Suit: deck.Club, Rank: deck.Ace},
			),
			normal: 11,
			soft:   1,
		},
		{
			hand: hand.New(
				deck.Card{Suit: deck.Club, Rank: deck.Jack},
				deck.Card{Suit: deck.Club, Rank: deck.Queen},
				deck.Card{Suit: deck.Spade, Rank: deck.King},
			),
			normal: 30,
			soft:   30,
		},
		{
			hand: hand.New(
				deck.Card{Suit: deck.Club, Rank: deck.Two},
				deck.Card{Suit: deck.Club, Rank: deck.Three},
				deck.Card{Suit: deck.Club, Rank: deck.Four},
				deck.Card{Suit: deck.Club, Rank: deck.Five},
			),
			normal: 14,
			soft:   14,
		},
		{
			hand: hand.New(
				deck.Card{Suit: deck.Club, Rank: deck.Six},
				deck.Card{Suit: deck.Club, Rank: deck.Seven},
				deck.Card{Suit: deck.Club, Rank: deck.Eight},
				deck.Card{Suit: deck.Club, Rank: deck.Nine},
				deck.Card{Suit: deck.Club, Rank: deck.Ten},
			),
			normal: 40,
			soft:   40,
		},
	}
	for _, test := range tables {
		normal, soft := test.hand.GetScores()
		if normal != test.normal || soft != test.soft {
			t.Fatalf("expected score: %d (%d), got: %d (%d)", test.normal, test.soft, normal, soft)
		}
	}
}

func TestHandScoreWithDraw(t *testing.T) {
	tables := []struct {
		hand         hand.Hand
		normal, soft int
	}{
		{
			hand:   hand.Hand{Cards: nil},
			normal: 0,
			soft:   0,
		},
		{
			hand:   hand.New(),
			normal: 0,
			soft:   0,
		},
		{
			hand: hand.Hand{
				Cards: &[]deck.Card{
					{Suit: deck.Club, Rank: deck.Six},
					{Suit: deck.Club, Rank: deck.Seven},
					{Suit: deck.Club, Rank: deck.Eight},
				},
			},
			normal: 21,
			soft:   21,
		},
		{
			hand: hand.New(
				deck.Card{Suit: deck.Club, Rank: deck.Two},
				deck.Card{Suit: deck.Heart, Rank: deck.Ace},
			),
			normal: 13,
			soft:   3,
		},
	}
	for _, test := range tables {
		normal, soft := test.hand.GetScores()
		if normal != test.normal || soft != test.soft {
			t.Fatalf("expected score before draw: %d (%d), got: %d (%d)", test.normal, test.soft, normal, soft)
		}

		test.hand.Draw(deck.Card{Suit: deck.Club, Rank: deck.Two}, deck.Card{Suit: deck.Spade, Rank: deck.Nine})
		lastDraw := deck.Card{Suit: deck.Heart, Rank: deck.Ace}
		test.hand.Draw(lastDraw)

		normal, soft = test.hand.GetScores()
		normalAfterDraw, softAfterDraw := test.normal+22, test.soft+12
		if normal != normalAfterDraw || soft != softAfterDraw {
			t.Fatalf("expected score after draw: %d (%d), got: %d (%d)", normalAfterDraw, softAfterDraw, normal, soft)
		}

		top := test.hand.TopCard()
		if top.Suit != lastDraw.Suit || top.Rank != lastDraw.Rank {
			t.Fatalf("expected top card: %s, got: %s", lastDraw, top)
		}
	}
}
