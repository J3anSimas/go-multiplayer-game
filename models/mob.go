package models

import "github.com/J3anSimas/game_multiplayer_go/types"

type Mob struct {
	Health      int
	Position    types.Point
	Strength    int
	CoinsToDrop int
}
