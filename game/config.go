package game

import (
	"crazy-eights/deck"
	"encoding/json"
	"io/ioutil"
	"os"
)

type ScoresConfig struct {
	Ace   int
	Two   int
	Three int
	Four  int
	Five  int
	Six   int
	Seven int
	Eight int
	Nine  int
	Ten   int
	Jack  int
	Queen int
	King  int
}

type Config struct {
	Scores ScoresConfig
}

func ReadConfig(fp string) (Config, error) {
	jsonFile, err := os.Open(fp)
	if err != nil {
		return Config{}, err
	}
	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	var config Config
	json.Unmarshal(byteValue, &config)

	return config, nil
}

func BuildScoreMap(c Config) map[deck.Rank]int {
	return map[deck.Rank]int{
		deck.Ace:   c.Scores.Ace,
		deck.Two:   c.Scores.Two,
		deck.Three: c.Scores.Three,
		deck.Four:  c.Scores.Four,
		deck.Five:  c.Scores.Five,
		deck.Six:   c.Scores.Six,
		deck.Seven: c.Scores.Seven,
		deck.Eight: c.Scores.Eight,
		deck.Nine:  c.Scores.Nine,
		deck.Ten:   c.Scores.Ten,
		deck.Jack:  c.Scores.Jack,
		deck.Queen: c.Scores.Queen,
		deck.King:  c.Scores.King,
	}
}
