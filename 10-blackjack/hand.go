package main

import (
	"fmt"
	"github.com/artslob/gophercises/09-deck/deck"
)

type hand []deck.Card

func (h *hand) draw(card deck.Card) {
	*h = append(*h, card)
}

func (h hand) score(soft bool) (score int) {
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
func (h hand) getScores() (int, int) {
	return h.score(false), h.score(true)
}

func (h hand) scoreString() string {
	normal, soft := h.getScores()
	return h.stringifyScores(normal, soft)
}

func (h hand) stringifyScores(normal, soft int) string {
	return fmt.Sprintf("%d (%d)", normal, soft)
}
