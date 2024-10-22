package models

import (
	"errors"

	"github.com/J3anSimas/game_multiplayer_go/types"
)

type GameLogic struct {
	Room *Room
}

func (g *GameLogic) MovePlayer(playerId string, dx, dy int) ([]types.Point, error) {
	player := g.Room.FindPlayerById(playerId)
	if player == nil {
		return nil, errors.New("jogador não encontrado")
	}
	return player.Move(dx, dy, g.Room)
}
