package deck

import "fmt"

type Rank int
type Suit int

const (
	Ace Rank = iota
	Two
	Three
	Four
	Five
	Six
	Seven
	Eight
	Nine
	Ten
	Jack
	Queen
	King
)

const (
	Diamond Suit = iota
	Club
	Heart
	Spade
)

type Card struct {
	R Rank
	S Suit
}

var Suits = [4]Suit{Diamond, Club, Heart, Spade}
var Ranks = [13]Rank{Ace, Two, Three, Four, Five, Six, Seven, Eight, Nine, Ten, Jack, Queen, King}

var SuitString = map[Suit]string{
	Diamond: "\u2666",
	Club:    "\u2663",
	Heart:   "\u2665",
	Spade:   "\u2660",
}

var RankString = map[Rank]string{
	Ace:   "A",
	Two:   "2",
	Three: "3",
	Four:  "4",
	Five:  "5",
	Six:   "6",
	Seven: "7",
	Eight: "8",
	Nine:  "9",
	Ten:   "10",
	Jack:  "J",
	Queen: "Q",
	King:  "K",
}

func (c Card) Stringify() string {
	return fmt.Sprintf("%v%v", RankString[c.R], SuitString[c.S])
}
