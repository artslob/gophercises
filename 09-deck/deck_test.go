package main

import (
	"github.com/artslob/gophercises/09-deck/deck"
	"testing"
)

const deckLength = 52

func TestDeckCreation(t *testing.T) {
	d := deck.New()
	if len(*d) != deckLength {
		t.Fatalf("expected deck contains %d elements, got: %d\n", deckLength, len(*d))
	}
	d.SortDefault()
	top := d.GetTopCard()
	if len(*d) != deckLength-1 {
		t.Fatal("expected length of deck decreased by 1")
	}
	if top.Suit != deck.Club {
		t.Fatal("top card`s suit is not Club but", top.Suit)
	}
	if top.Rank != deck.King {
		t.Fatal("top card`s rank is not King but", top.Rank)
	}
}

func TestCustomSorting(t *testing.T) {
	d := deck.New()
	d.Sort(func(i, j deck.Card) bool {
		return i.Rank > j.Rank
	})
	top := d.GetTopCard()
	if len(*d) != deckLength-1 {
		t.Fatal("expected length of deck decreased by 1")
	}
	if top.Suit != deck.Club {
		t.Fatal("top card`s suit is not Club but", top.Suit)
	}
	if top.Rank != deck.Ace {
		t.Fatal("top card`s rank is not Ace but", top.Rank)
	}
}

func TestGettingTopCard(t *testing.T) {
	d := deck.New()
	d.SortDefault()
	if len(*d) != deckLength {
		t.Fatalf("expected deck contains %d elements, got: %d\n", deckLength, len(*d))
	}
	times := deckLength / 2
	for i := 0; i < times; i++ {
		d.GetTopCard()
	}
	if len(*d) != times {
		t.Fatalf("expected deck contains %d elements, got: %d\n", times, len(*d))
	}
	for i := 0; i < times; i++ {
		d.GetTopCard()
	}
	if len(*d) != 0 {
		t.Fatalf("expected deck contains 0 elements, got: %d\n", len(*d))
	}
}
