package main

import (
	"crazy-eights/game"
	"fmt"
)

func main() {
	g := game.NewGame(4)
	g.Deal(5)
	for _, hand := range g.ShowHands() {
		fmt.Println(hand)
	}
}
