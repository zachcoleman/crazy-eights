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

	winnerMap := make(map[int]int)

	for i := 0; i < 1_000; i++ {
		// create and setup game
		g := game.NewGame(2)
		g.Deal(5)
		g.PullStarter()

		// play game based on scoring configuration
		g.Play(scoreMap)

		if _, ok := winnerMap[g.Winner]; ok {
			winnerMap[g.Winner] += 1
		} else {
			winnerMap[g.Winner] = 1
		}
	}
	fmt.Println(winnerMap)

}
