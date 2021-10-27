package main

import (
	"crazy-eights/game"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	g := game.NewGame(4)
	g.Deal(5)
	g.Play()
}
