package game

type Strategy interface {
	PickMove([]Move) Move
}
