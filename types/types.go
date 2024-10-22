package types

type Status uint8

type Point struct {
	X, Y int
}

const (
	WaitingForConnection Status = iota
	Running
	Paused
	GameOver
)
