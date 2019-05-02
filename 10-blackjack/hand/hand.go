package hand

import (
	"fmt"
	"github.com/artslob/gophercises/09-deck/deck"
)

type Hand []deck.Card

func (h *Hand) Draw(card deck.Card) {
	*h = append(*h, card)
}

func (h Hand) Score(soft bool) (score int) {
	for _, card := range h {
		switch card.Rank {
		case deck.Ace:
			if soft {
				score++
			} else {
				score += 11
			}
		case deck.Jack, deck.Queen, deck.King:
			score += 10
		default:
			score += int(card.Rank) + 1
		}
	}
	return
}

// getScores first result - is normal score; second - is soft (when Ace equals to 1)
func (h Hand) GetScores() (int, int) {
	return h.Score(false), h.Score(true)
}

func (h Hand) ScoreString() string {
	normal, soft := h.GetScores()
	return h.StringifyScores(normal, soft)
}

func (h Hand) StringifyScores(normal, soft int) string {
	return fmt.Sprintf("%d (%d)", normal, soft)
}
