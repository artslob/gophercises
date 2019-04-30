package main

import (
	"fmt"
	"github.com/artslob/gophercises/09-deck/deck"
)

func main() {
	d := deck.New()
	d.Shuffle()
	d.Sort(func(i, j deck.Card) bool {
		return i.Suit > j.Suit
	})
	d.SortDefault()
	for _, card := range *d {
		fmt.Println(card)
	}
}
