package hand

import (
	"fmt"
	"github.com/artslob/gophercises/09-deck/deck"
)

const Blackjack = 21

type Hand struct {
	Cards         *[]deck.Card
	normalScore   int
	softScore     int
	lastCalcIndex int
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

func (h *Hand) Score() int {
	h.calcScore()
	return h.normalScore
}

func (h *Hand) SoftScore() int {
	h.calcScore()
	return h.softScore
}

// getScores first result - is normal score; second - is soft (when Ace equals to 1)
func (h *Hand) GetScores() (int, int) {
	h.calcScore()
	return h.Score(), h.SoftScore()
}

func (h *Hand) ScoreString() string {
	normal, soft := h.GetScores()
	return fmt.Sprintf("%d (%d)", normal, soft)
}

func (h *Hand) calcScore() {
	if h == nil {
		return
	}
	if h.Cards == nil || *h.Cards == nil {
		h.normalScore, h.softScore = 0, 0
		return
	}
	for ; h.lastCalcIndex < len(*h.Cards); h.lastCalcIndex++ {
		card := (*h.Cards)[h.lastCalcIndex]
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
