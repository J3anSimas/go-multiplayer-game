/**
*
	@type{number} - The status of the game
* 
* */

c

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
WaitingForPlayersToGetReady
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
