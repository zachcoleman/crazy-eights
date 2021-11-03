package game

import (
	"crazy-eights/deck"
	"fmt"
	"math"
)

type Strategy interface {
	PickMove(g Game) Move
}

type SimpleStrategy struct{}
type EvalStratgy struct {
	ScoreMap map[deck.Rank]int
}
type MinimaxStrategy struct {
	ScoreMap map[deck.Rank]int
	MaxDepth int
}

func (s SimpleStrategy) PickMove(g *Game) Move {
	moves := g.GetMoves()
	tmp := map[string]Move{}

	for _, m := range moves {
		if IsPass(m) {
			tmp["pass"] = m
		} else {
			tmp["move"] = m
			break
		}
	}

	if m, ok := tmp["move"]; ok {
		return m
	}
	if m, ok := tmp["pass"]; ok {
		return m
	}
	return moves[0]
}

func (s EvalStratgy) PickMove(g *Game) Move {
	moves := g.GetMoves()
	evals := []int{}

	for _, m := range moves {
		tmpGame := g.CopyGame()
		tmpGame.PlayMove(m)
		evals = append(evals, tmpGame.Eval(s.ScoreMap))
	}

	// for idx, _ := range moves {
	// 	fmt.Printf(
	// 		"Playing %v, results in eval: %v \n",
	// 		moves[idx].Card.Stringify(),
	// 		evals[idx],
	// 	)
	// }

	idx, _ := min(evals...)
	return moves[idx]
}

func (s MinimaxStrategy) PickMove(g *Game) Move {
	_, move := minimax(
		g.CopyGame(),
		g.GetCurrentPlayer().copyPlayer(),
		s.ScoreMap,
		s.MaxDepth,
		0,
	)
	return move
}

// minimax for the target player this is for finding highest scoring evaluation
func minimax(g Game, p Player, scoreMap map[deck.Rank]int, maxDepth int, currentDepth int) (int, Move) {

	if g.IsGameOver() || currentDepth == maxDepth {
		return Eval(g.CopyGame(), p.copyPlayer(), scoreMap), PassMove(p.ID)
	}

	var bestScore int
	var bestMove Move

	if g.GetCurrentPlayerID() == p.ID {
		bestScore = math.MinInt64
	} else {
		bestScore = math.MaxInt64
	}

	for _, move := range g.GetMoves() {

		// copy game, play it, and advance game
		newGame := g.CopyGame()
		newGame.PlayMove(move)
		newGame.NextPlayer()

		// recursive step
		currentScore, _ := minimax(newGame, p, scoreMap, maxDepth, currentDepth+1)

		if currentDepth == 0 {
			fmt.Printf(
				"player: %v, move: %v @ score: %v \n",
				g.GetCurrentPlayerID(),
				move.Stringify(),
				currentScore,
			)
		}

		// if its target player then take highest score
		// else take minimizing score
		if g.GetCurrentPlayerID() == p.ID {
			if currentScore > bestScore {
				bestScore = currentScore
				bestMove = move
			}
		} else {
			if currentScore < bestScore {
				bestScore = currentScore
				bestMove = move
			}
		}
	}

	return bestScore, bestMove
}
