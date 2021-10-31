package game

type Strategy interface {
	PickMove(g Game) Move
}

type SimpleStrategy struct{}

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
