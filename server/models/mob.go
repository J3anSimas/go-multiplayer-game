package models

import "github.com/J3anSimas/game_multiplayer_go/types"

type Mob struct {
	Health      int
	Position    types.Point
	Strength    int
	CoinsToDrop int
}

func NewMob(health int, position types.Point, strength int, coins int) *Mob {
	return &Mob{
		Health:      health,
		Position:    position,
		Strength:    strength,
		CoinsToDrop: coins,
	}
}
