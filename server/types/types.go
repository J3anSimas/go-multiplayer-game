package types

type Status uint8

type Point struct {
	X, Y int
}

type Turn uint8

const (
	HostTurn Turn = iota
	GuestTurn
)
const (
	WaitingForGuestConnection Status = iota
	WaitingForReady
	Running
	Paused
	GameOver
)

type ShopItemAttributeModifier uint8

const (
	StrengthAttribute ShopItemAttributeModifier = iota
	MovementAttribute
	AttackVelocityAttribute
)
