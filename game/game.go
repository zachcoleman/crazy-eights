package game

import (
	"container/ring"
	"crazy-eights/deck"
	"fmt"
	"strings"
)

const Wild deck.Rank = deck.Eight

type Game struct {
	Players     *ring.Ring
	DiscardPile deck.Deck
	DrawPile    deck.Deck
	Dealt       bool
	WildRank    deck.Rank
	Over        bool
	Winner      int
}

func (g *Game) GetCurrentPlayer() *Player       { return g.Players.Value.(*Player) }
func (g *Game) GetCurrentPlayerHand() deck.Deck { return deck.Deck(g.Players.Value.(*Player).Hand) }
func (g *Game) GetCurrentPlayerID() int         { return g.Players.Value.(*Player).ID }
func (g *Game) NextPlayer()                     { g.Players = g.Players.Next() }

func NewGame(numPlayers int) Game {
	// create players
	players := ring.New(numPlayers)
	for i := 0; i < numPlayers; i++ {
		tmp := NewPlayer(i)
		players.Value = &tmp
		players = players.Next()
	}

	// create DrawPile and shuffle it
	tmp := deck.NewFullDeck()
	tmp.Shuffle()

	// build game
	return Game{
		Players:     players,
		DiscardPile: []deck.Card{},
		DrawPile:    tmp,
		Dealt:       false,
		WildRank:    Wild,
		Winner:      -1,
	}
}

func (g *Game) CopyGame() Game {
	// create new same-sized game
	gameCopy := NewGame(g.Players.Len())

	// copy players
	for i := 0; i < gameCopy.Players.Len(); i++ {
		currPlayer := g.GetCurrentPlayer()
		newPlayer := currPlayer.copyPlayer()
		gameCopy.Players.Value = &newPlayer
		gameCopy.Players = gameCopy.Players.Next()
		g.Players = g.Players.Next()
	}

	// copy other fields
	gameCopy.DiscardPile = g.DiscardPile.CopyDeck()
	gameCopy.DrawPile = g.DrawPile.CopyDeck()
	gameCopy.Dealt = g.Dealt
	gameCopy.WildRank = g.WildRank
	gameCopy.Winner = g.Winner

	return gameCopy
}

func (g *Game) Deal(numCards int) {
	for j := 0; j < numCards; j++ {
		for i := 0; i < g.Players.Len(); i++ {
			c, err := g.DrawPile.Draw()
			if err != nil {
				fmt.Println("failed to draw card")
			}
			g.Players.Value.(*Player).AddToHand(c)
			g.Players = g.Players.Next()
		}
	}
	g.Dealt = true
}

func (g *Game) PullStarter() {
	c, _ := g.DrawPile.Draw()
	g.DiscardPile = append(g.DiscardPile, c)
}

func (g Game) ShowHands() []string {
	hands := []string{}
	for i := 0; i < g.Players.Len(); i++ {
		tmp := []string{}
		for _, c := range g.Players.Value.(*Player).Hand {
			tmp = append(tmp, c.Stringify())
		}
		id := g.Players.Value.(*Player).ID
		hands = append(hands,
			fmt.Sprintf("Player %v: ", id)+strings.Join(tmp, ","),
		)
		g.Players = g.Players.Next()
	}
	return hands
}

func (g Game) validCard(c deck.Card) bool {
	if len(g.DiscardPile) == 0 {
		return true
	}

	currentDiscard := g.DiscardPile[len(g.DiscardPile)-1]
	if c.R == currentDiscard.R {
		return true
	} else if c.S == currentDiscard.S {
		return true
	} else if c.R == g.WildRank {
		return true
	}
	return false
}

func (g Game) ValidMove(m Move) bool {
	if IsPass(m) && len(g.DrawPile) > 0 {
		return true
	}
	return (g.GetCurrentPlayerID() == m.PlayerId && // correct player
		g.validCard(m.Card) && // valid card
		m.Card == g.GetCurrentPlayerHand()[m.cardIdx]) // correct index
}

func (g *Game) GetMoves() []Move {
	moves := []Move{}
	id := g.GetCurrentPlayerID()

	// build and check pass move
	passMove := PassMove(id)
	if g.ValidMove(passMove) {
		moves = append(moves, passMove)
	}

	// build and check moves from cards
	hand := g.GetCurrentPlayerHand()
	for cardIdx, card := range hand {
		m := NewMove(id, card, cardIdx)
		if g.ValidMove(m) {
			moves = append(moves, m)
		}
	}

	return moves
}

func (g *Game) PlayMove(m Move) error {
	if g.ValidMove(m) {
		currPlayerPtr := g.GetCurrentPlayer()
		currPlayerHand := g.GetCurrentPlayerHand()

		if IsPass(m) {
			// draw card
			c, err := g.DrawPile.Draw()
			if err != nil {
				return err
			}
			currPlayerHand = append(currPlayerHand, c)
		} else {
			// play card
			g.DiscardPile = append(g.DiscardPile, m.Card)
			currPlayerHand[m.cardIdx] = currPlayerHand[len(currPlayerHand)-1]
			currPlayerHand = currPlayerHand[:len(currPlayerHand)-1]
		}

		// update hand
		(*currPlayerPtr).Hand = currPlayerHand

		// no error to return
		return nil

	} else {
		// not valid turn
		return fmt.Errorf("%v not a valid turn", m.Stringify())
	}
}

// TODO: should this check every player or just current one
func (g Game) IsGameOver() bool {
	currPlayerHand := g.GetCurrentPlayerHand()
	if len(currPlayerHand) == 0 {
		return true
	}
	if len(g.DrawPile) == 0 {
		return true
	}
	return false
}

func (g Game) Score(scoreMap map[deck.Rank]int) map[int]int {
	scores := map[int]int{}
	for i := 0; i < g.Players.Len(); i++ {
		p := g.GetCurrentPlayer()
		scores[p.ID] = p.Score(scoreMap)
		g.Players = g.Players.Next()
	}
	return scores
}

func (g *Game) MarkWinner(scores map[int]int) {
	id, minscore := 0, scores[0]
	for i, val := range scores {
		if val < minscore {
			minscore = val
			id = i
		}
	}
	g.Winner = id
}

func (g *Game) Play(scoreMap map[deck.Rank]int) {
	// game loop
	for {
		// DEBUG: print player hands
		fmt.Println("====== TURN ======")
		fmt.Printf("Current Player ID: %v \n", g.GetCurrentPlayerID())
		fmt.Printf("Current Player Hand: %v \n", g.GetCurrentPlayerHand().Stringify())

		// DEBUG: print the discard pile
		if len(g.DiscardPile) == 1 {
			fmt.Println(
				"Discard:",
				g.DiscardPile[len(g.DiscardPile)-1].Stringify(),
			)
		} else if len(g.DiscardPile) > 1 {
			fmt.Println(
				"Discard:",
				g.DiscardPile[len(g.DiscardPile)-1].Stringify(),
				g.DiscardPile[len(g.DiscardPile)-2].Stringify(),
			)
		}

		// ====== implement strategies ======
		m := Move{}
		if g.GetCurrentPlayerID() == 0 {
			m = SimpleStrategy{}.PickMove(g)
		} else {
			m = EvalStratgy{ScoreMap: scoreMap}.PickMove(g)
		}
		fmt.Printf("Player: %v, card: %v \n", m.PlayerId, m.Card.Stringify())
		g.PlayMove(m)
		// ======================

		scores := g.Score(scoreMap)
		for id, score := range scores {
			fmt.Printf("Player %v: %v \n", id, score)
		}

		if g.IsGameOver() {
			g.MarkWinner(scores)
			break
		}

		g.NextPlayer()

	}

	fmt.Println("====== END OF GAME ======")
	if g.Winner != -1 {
		fmt.Printf("!!! Winner is Player %v !!!\n", g.Winner)
	}
}
