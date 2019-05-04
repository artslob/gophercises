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
	h.Draw(cards...)
	return h
}

func (h *Hand) Draw(cards ...deck.Card) {
	if h.Cards == nil {
		h.Cards = &[]deck.Card{}
	}
	*h.Cards = append(*h.Cards, cards...)
	h.calcScore()
}

func (h Hand) TopCard() deck.Card {
	return (*h.Cards)[len(*h.Cards)-1]
}

func (h Hand) Size() int {
	return len(*h.Cards)
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

func (h *Hand) BestScore() int {
	h.calcScore()
	if h.normalScore <= Blackjack {
		return h.normalScore
	}
	return h.softScore
}

func (h *Hand) IsBlackjack() bool {
	h.calcScore()
	return h.softScore == Blackjack || h.normalScore == Blackjack
}

func (h *Hand) MoreThanBlackjack() bool {
	h.calcScore()
	return h.softScore > Blackjack && h.normalScore > Blackjack
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
