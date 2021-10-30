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
	DiscardPile []deck.Card
	DrawPile    deck.Deck
	Dealt       bool
	WildRank    deck.Rank
	Winner      int
}

func (g *Game) GetCurrentPlayer() *Player {
	return g.Players.Value.(*Player)
}

func (g *Game) GetCurrentPlayerHand() deck.Deck {
	return deck.Deck(g.Players.Value.(*Player).Hand)
}

func (g *Game) GetCurrentPlayerID() int {
	return g.Players.Value.(*Player).ID
}

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

// TODO: decide if this should be pointer or value reciever
// both have merit -> value reciever makes it extra safe on not
// mutating the orig. game being copied
func (g *Game) CopyGame() Game {
	// create new same-sized game
	gameCopy := NewGame(g.Players.Len())

	// copy players
	for i := 0; i < gameCopy.Players.Len(); i++ {
		tmp := g.Players.Value.(*Player).copyPlayer()
		gameCopy.Players.Value = &tmp
		gameCopy.Players = gameCopy.Players.Next()
		g.Players = g.Players.Next()
	}

	// copy other fields
	gameCopy.DiscardPile = g.DiscardPile
	gameCopy.DrawPile = g.DrawPile
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

func (g *Game) ValidCard(c deck.Card) bool {
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

func (g *Game) Play(scoreMap map[deck.Rank]int) {

	// status
	// skipped := map[int]bool{}
	// for i := 0; i < g.Players.Len(); i++ {
	// 	skipped[g.GetCurrentPlayerID()] = false
	// 	g.Players = g.Players.Next()
	// }

	// game loop
	for {
		fmt.Println("====== TURN ======")
		for _, hand := range g.ShowHands() {
			fmt.Println(hand)
		}
		fmt.Printf("Current Player ID: %v \n", g.GetCurrentPlayerID())
		fmt.Printf("Current Player Hand: %v \n", g.GetCurrentPlayerHand().Stringify())
		fmt.Printf("Eval: %v \n", g.Eval(scoreMap))

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

		// get current player ptr and hand
		currPlayerPtr := g.GetCurrentPlayer()
		currPlayerHand := g.GetCurrentPlayerHand()
		played := false

		// for card in hand
		for idx, c := range currPlayerHand {

			// empty discard so play whatever
			if len(g.DiscardPile) == 0 {
				g.DiscardPile = append(g.DiscardPile, c)
				currPlayerHand[idx] = currPlayerHand[len(currPlayerHand)-1]
				currPlayerHand = currPlayerHand[:len(currPlayerHand)-1]
				played = true
				break
			}

			// once find valid card play it
			if g.ValidCard(c) {
				g.DiscardPile = append(g.DiscardPile, c)
				currPlayerHand[idx] = currPlayerHand[len(currPlayerHand)-1]
				currPlayerHand = currPlayerHand[:len(currPlayerHand)-1]
				played = true
				break
			}
		}

		// if no card played then draw
		if !played {
			c, _ := g.DrawPile.Draw()
			currPlayerHand = append(currPlayerHand, c)
		}

		// update skipped
		// skipped[currPlayerPtr.ID] = played
		// if checkAllTrue(skipped) {
		// 	break
		// }

		// update hand
		(*currPlayerPtr).Hand = currPlayerHand

		// if out of cards finish game
		if len(currPlayerHand) == 0 {
			g.Winner = currPlayerPtr.ID
			break
		}

		if len(g.DrawPile) == 0 {
			// TODO calculate scores
			break
		}

		// go to next player
		g.Players = g.Players.Next()
	}

	fmt.Println("====== END OF GAME ======")
	if g.Winner != -1 {
		fmt.Printf("!!! Winner is Player %v !!!\n", g.Winner)
	}
}

func checkAllTrue(check map[int]bool) bool {
	for _, val := range check {
		if !val {
			return false
		}
	}
	return true
}
