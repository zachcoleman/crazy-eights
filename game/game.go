package game

import (
	"container/ring"
	"crazy-eights/deck"
	"fmt"
	"strings"
)

type Player struct {
	Hand []deck.Card
}

type Game struct {
	Players     *ring.Ring
	DiscardPile []deck.Card
	DrawPile    deck.Deck
	Dealt       bool
}

func NewGame(numPlayers int) Game {
	players := ring.New(numPlayers)
	for i := 0; i < numPlayers; i++ {
		players.Value = &Player{
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
		hands = append(hands, strings.Join(tmp, ","))
		g.Players = g.Players.Next()
	}
	return hands
}
