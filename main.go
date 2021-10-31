package main

import (
	"crazy-eights/game"
	"math/rand"
	"time"
)

func main() {
	// seed
	rand.Seed(time.Now().Unix())

	// read in scoring configuration
	config, err := game.ReadConfig("config.json")
	if err != nil {
		panic(err)
	}
	scoreMap := game.BuildScoreMap(config)

	// create and setup game
	g := game.NewGame(2)
	g.Deal(5)
	g.PullStarter()

	// play game based on scoring configuration
	g.Play(scoreMap)
}
