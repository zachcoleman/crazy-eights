package game

import (
	"crazy-eights/deck"
)

// Eval score the game state for the currently pointed
// at player based on the scoreMap
func (g *Game) Eval(scoreMap map[deck.Rank]int) int {
	loss := 0
	hand := g.GetCurrentPlayerHand()
	for _, card := range hand {
		score := scoreMap[card.R]
		if g.ValidCard(card) {
			loss += score / 2
		} else {
			loss += score
		}
	}
	return loss
}
