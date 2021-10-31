package game

import (
	"crazy-eights/deck"
)

type Strategy interface {
	PickMove(g Game) Move
}

type SimpleStrategy struct{}
type EvalStratgy struct {
	ScoreMap map[deck.Rank]int
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
