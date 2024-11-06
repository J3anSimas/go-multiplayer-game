declare enum Turn {
	HostTurn,
	GuestTurn,
}
declare enum Status {
	WaitingForGuestConnection,
	WaitingForPlayersToGetReady,
	Running,
	Paused,
	GameOver,
}

declare enum ShopItemAttributeModifier {
	StrengthAttribute,
	MovementAttribute,
	AttackVelocityAttribute,
}
type Player = {
	Id: string;
	Ready: boolean;
	IsHost: boolean;
	Position: Point;
	Health: number;
	Coins: number;
	MoveCapacity: number;
	MovesRemaining: number;
	Strength: number;
	TotalShots: number;
	ShotsRemaining: number;
	IsDead: boolean;
};
type Point = {
	X: number;
	Y: number;
};

type Mob = {
	Health: number;
	Position: Point;
	Strength: number;
	CoinsToDrop: number;
};

type GameState = {
	Id: string;
	State: Status;
	WorldWidth: number;
	WorldHeight: number;
	Players: Player[];
	Mobs: Mob[];
	ShopItems: ShopItem[];
	Turn: Turn;
};

type ShopItem = {
	Title: string;
	Description: string;
	Cost: number;
	Attribute: ShopItemAttributeModifier;
	Modifier: number;
};

