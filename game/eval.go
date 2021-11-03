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
		loss += score
	}
	return loss
}

func Eval(g Game, p Player, scoreMap map[deck.Rank]int) int {
	loss := 0
	g.GoToPlayer(p.ID)
	hand := g.GetCurrentPlayerHand()
	for _, card := range hand {
		score := scoreMap[card.R]
		loss += score
	}
	return -1 * loss
}
