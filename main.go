package main

import (
	"crazy-eights/game"
	"fmt"
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

	numPlayers := 2
	winnerMap := make(map[int]int)
	for i := 0; i < numPlayers; i++ {
		winnerMap[i] = 0
	}

	for i := 0; i < 100; i++ {
		// create and setup game
		g := game.NewGame(numPlayers)
		g.Deal(5)
		g.PullStarter()
		g.GoToPlayer(i % numPlayers) // alternate who plays first

		// play game based on scoring configuration
		g.Play(scoreMap)
		winnerMap[g.Winner] += 1
	}
	fmt.Println(winnerMap)

}
