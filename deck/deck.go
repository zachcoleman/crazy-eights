package deck

import (
	"fmt"
	"math/rand"
	"strings"
)

// Deck will refer to any slice of cards
type Deck []Card

// NewFullDeck creates a new standard 52 card
// Deck of Cards
func NewFullDeck() Deck {
	ret := Deck{}
	for _, s := range Suits {
		for _, r := range Ranks {
			ret = append(ret, Card{r, s})
		}
	}
	return ret
}

func (d Deck) Shuffle() {
	for i := 0; i < len(d); i++ {
		x, y := rand.Intn(len(d)), rand.Intn(len(d))
		d[x], d[y] = d[y], d[x]
	}
}

func (d *Deck) Draw() (Card, error) {
	if len(*d) > 0 {
		var ret Card
		ret, *d = (*d)[0], (*d)[1:]
		return ret, nil
	}
	return Card{-1, 1}, fmt.Errorf("attempting to draw from empty deck")
}

func (d Deck) Stringify() string {
	tmp := []string{}
	for _, c := range d {
		tmp = append(tmp, c.Stringify())
	}
	return strings.Join(tmp, ",")
}

func (d Deck) ValidateDeck() bool {
	cards := make([][]int, len(Suits))
	for idx := range cards {
		cards[idx] = make([]int, len(Ranks))
	}

	for _, card := range d {
		if int(card.S) < len(cards) && int(card.R) <= len(cards[0]) && 0 <= int(card.S) && 0 <= int(card.R) {
			cards[card.S][card.R] += 1
		} else {
			return false
		}
	}

	sum := 0
	for _, ranks := range cards {
		for _, val := range ranks {
			sum += val
		}
	}

	return sum == len(Suits)*len(Ranks)
}
