package deck

import "fmt"

type Suit byte

const (
	// ♠
	Spade Suit = iota
	// ♥
	Heart
	// ♦
	Diamond
	// ♣
	Club
)

type Rank byte

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

type Card struct {
	Suit Suit
	Rank Rank
}

func (c Card) String() string {
	return fmt.Sprintf("%s: %s", c.Suit, c.Rank)
}
