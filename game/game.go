package game

import (
	"container/ring"
	"crazy-eights/deck"
	"fmt"
	"strings"
)

type Player struct {
	ID   int
	Hand []deck.Card
}

type Game struct {
	Players     *ring.Ring
	DiscardPile []deck.Card
	DrawPile    deck.Deck
	Dealt       bool
	WildRank    deck.Rank
	Winner      int
}

func NewGame(numPlayers int) Game {
	players := ring.New(numPlayers)
	for i := 0; i < numPlayers; i++ {
		players.Value = &Player{
			ID:   i,
			Hand: []deck.Card{},
		}
		players = players.Next()
	}
	tmp := deck.CreateNewDeck()
	tmp.Shuffle()

	return Game{
		Players:     players,
		DiscardPile: []deck.Card{},
		DrawPile:    tmp,
		Dealt:       false,
		WildRank:    deck.Eight,
		Winner:      -1,
	}
}

func (g Game) Deal(numCards int) {
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

func (p *Player) AddToHand(c deck.Card) {
	p.Hand = append(p.Hand, c)
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

func (g Game) ValidCard(c deck.Card) bool {
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

func (g Game) Play() {
	// count := 0

	// game loop
	for {
		fmt.Println("====== TURN ======")
		for _, hand := range g.ShowHands() {
			fmt.Println(hand)
		}
		if len(g.DiscardPile) == 1 {
			fmt.Println(g.DiscardPile[len(g.DiscardPile)-1].Stringify())
		} else if len(g.DiscardPile) > 1 {
			fmt.Println(g.DiscardPile[len(g.DiscardPile)-1].Stringify(), g.DiscardPile[len(g.DiscardPile)-2].Stringify())
		}

		currPlayerPtr := g.Players.Value.(*Player)
		currPlayerHand := (*currPlayerPtr).Hand
		played := false

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

		// update hand
		(*currPlayerPtr).Hand = currPlayerHand

		// if out of cards finish game
		if len(currPlayerHand) == 0 {
			g.Winner = currPlayerPtr.ID
			break
		}

		if len(g.DrawPile) == 0 {
			// TODO calculate scores
		}

		// go to next player
		g.Players = g.Players.Next()

		// if count > 10 {
		// 	break
		// }
		// count++
	}

	fmt.Println("====== END OF GAME ======")
	if g.Winner != -1 {
		fmt.Printf("!!! Winner is Player %v !!!\n", g.Winner)
	}
	for _, hand := range g.ShowHands() {
		fmt.Println(hand)
	}
}
