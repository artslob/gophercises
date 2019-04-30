package deck

import (
	"math/rand"
	"sort"
	"time"
)

type Deck []Card

func New() *Deck {
	result := Deck{}
	for suit := Spade; suit <= Club; suit++ {
		for rank := Ace; rank <= King; rank++ {
			result = append(result, Card{Suit: suit, Rank: rank})
		}
	}
	return &result
}

type defaultSortOrder Deck

func (d defaultSortOrder) Len() int {
	return len(d)
}

func (d defaultSortOrder) Less(i, j int) bool {
	if d[i].Suit < d[j].Suit {
		return true
	}
	if d[i].Suit == d[j].Suit {
		return d[i].Rank < d[j].Rank
	}
	return false
}

func (d defaultSortOrder) Swap(i, j int) {
	d[i], d[j] = d[j], d[i]
}

func (d Deck) SortDefault() {
	sort.Sort(defaultSortOrder(d))
}

func (d Deck) Sort(less func(i, j Card) bool) {
	sort.Slice(d, func(i, j int) bool {
		return less(d[i], d[j])
	})
}

func (d Deck) Shuffle() {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	r.Shuffle(len(d), defaultSortOrder(d).Swap)
}
