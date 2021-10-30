package main

import (
	"crazy-eights/game"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().Unix())
	g := game.NewGame(2)
	g.Deal(5)
	g.PullStarter()
	config, err := game.ReadConfig("config.json")
	if err != nil {
		panic(err)
	}
	scoreMap := game.BuildScoreMap(config)
	g.Play(scoreMap)
}
