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

func (g *GameLogic) PlayerAttackAnotherPlayer(attacker *Player, target *Player) error {
	if target.IsDead {
		return errors.New("Alvo já está morto")
	}
	if attacker.ShotsRemaining == 0 {
		return errors.New("Atacante não possui ataques restantes")
	}

	target.Health -= attacker.Strength
	attacker.ShotsRemaining--
	if target.Health <= 0 {
		target.IsDead = true
		g.FinishGame(*attacker)
	}
	return nil

}

func (g *GameLogic) PlayerAttackMob(attacker *Player, target *Mob) error {
	if attacker.ShotsRemaining == 0 {
		return errors.New("Atacante não possui ataques restantes")
	}

	target.Health -= attacker.Strength
	attacker.ShotsRemaining--
	if target.Health <= 0 {
		attacker.Coins += target.CoinsToDrop
		target = nil
	}
	return nil

}
func (g *GameLogic) FinishGame(winner Player) string {
	g.Room.Status = types.GameOver
	return winner.Id
}
