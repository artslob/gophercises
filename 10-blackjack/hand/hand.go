package hand

import (
	"fmt"
	"github.com/artslob/gophercises/09-deck/deck"
)

type Hand struct {
	Cards       *[]deck.Card
	normalScore int
	softScore   int
}

func New(cards ...deck.Card) Hand {
	h := Hand{Cards: &[]deck.Card{}}
	for _, card := range cards {
		h.Draw(card)
	}
	return h
}

func (h *Hand) Draw(card deck.Card) {
	if h.Cards == nil {
		h.Cards = &[]deck.Card{}
	}
	*h.Cards = append(*h.Cards, card)
	h.calcScore()
}

func (h Hand) Score() int {
	return h.normalScore
}

func (h Hand) SoftScore() int {
	return h.softScore
}

func (h *Hand) calcScore() {
	h.normalScore, h.softScore = 0, 0
	for _, card := range *h.Cards {
		switch card.Rank {
		case deck.Ace:
			h.normalScore += 11
			h.softScore++
		case deck.Jack, deck.Queen, deck.King:
			h.normalScore += 10
			h.softScore += 10
		default:
			h.normalScore += int(card.Rank) + 1
			h.softScore += int(card.Rank) + 1
		}
	}
}

// getScores first result - is normal score; second - is soft (when Ace equals to 1)
func (h Hand) GetScores() (int, int) {
	return h.normalScore, h.softScore
}

func (h Hand) ScoreString() string {
	normal, soft := h.GetScores()
	return h.StringifyScores(normal, soft)
}

func (h Hand) StringifyScores(normal, soft int) string {
	return fmt.Sprintf("%d (%d)", normal, soft)
}
