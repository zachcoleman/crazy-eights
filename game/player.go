package game

import (
	"crazy-eights/deck"
	"fmt"
)

type Player struct {
	ID   int
	Hand deck.Deck
}

func NewPlayer(id int) Player {
	return Player{
		ID:   id,
		Hand: []deck.Card{},
	}
}

func (p Player) copyPlayer() Player {
	return Player{
		p.ID,
		p.Hand.CopyDeck(),
	}
}

func (p *Player) AddToHand(c deck.Card) {
	p.Hand = append(p.Hand, c)
}

func (p Player) Stringify() string {
	return fmt.Sprintf(
		"Player ID: %v; Hand: %v",
		p.ID,
		deck.Deck(p.Hand).Stringify(),
	)
}

func (p Player) Score(scoreMap map[deck.Rank]int) int {
	score := 0
	for _, card := range p.Hand {
		score += scoreMap[card.R]
	}
	return score
}
