package models

import "github.com/J3anSimas/game_multiplayer_go/types" // Atualize com o nome do seu módulo

type Mob struct {
	Health   int
	Position types.Point
	Strength int
}
