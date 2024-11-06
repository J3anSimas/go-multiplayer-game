/** 
 * @readonly
 * @enum {number}
*/
export var Turn = {
	HostTurn: 0,
	GuestTurn: 1,
}
/**
 * @readonly
 * @enum {number}
 * */
export var State = {
	WaitingForGuestConnection: 0,
	WaitingForPlayersToGetReady: 1,
	Running: 2,
	Paused: 3,
	GameOver: 4,
}

/**
 * @readonly
 * @enum {number}
 * */
export var ShopItemAttributeModifier = {
	StrengthAttribute: 0,
	MovementAttribute: 1,
	AttackVelocityAttribute: 2,
}

