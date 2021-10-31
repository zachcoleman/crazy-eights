package game

import (
	"crazy-eights/deck"
	"fmt"
)

type Move struct {
	PlayerId int
	Card     deck.Card
	cardIdx  int
}

func NewMove(id int, card deck.Card, cardIdx int) Move {
	return Move{
		id,
		card,
		cardIdx,
	}
}

func PassMove(id int) Move {
	return Move{
		PlayerId: id,
		Card:     deck.Card{R: -1, S: -1},
		cardIdx:  -1,
	}
}

func IsPass(m Move) bool {
	noCard := deck.Card{R: -1, S: -1}
	if m.Card == noCard {
		return true
	}
	if m.cardIdx == -1 {
		return true
	}
	return false
}

func (m Move) Stringify() string {
	return fmt.Sprintf(
		"PlayerID: %v, Card: %v",
		m.PlayerId,
		m.Card.Stringify(),
	)
}
