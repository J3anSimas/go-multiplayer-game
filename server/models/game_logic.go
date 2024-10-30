package models

import (
	"errors"

	"github.com/J3anSimas/game_multiplayer_go/types"
)

type GameLogic struct {
	Room *Room
}

func (g *GameLogic) MovePlayer(player *Player,
	x, y int,
) ([]types.Point, error) {
	return player.Move(x, y, g.Room)
}

func (g *GameLogic) PlayerAttackAnotherPlayer(attacker *Player,
	target *Player) error {
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
func (g GameLogic) ListShopItems() []ShopItem {
	return g.Room.ShopItems
}
func (g *GameLogic) ChangeTurn(player *Player) {
	if player.IsHost {
		g.Room.Turn = types.GuestTurn
	} else {
		g.Room.Turn = types.HostTurn
	}
	player.ResetAttributes()
}

func (g GameLogic) BuyItem(player *Player, shopItem ShopItem) error {
	switch shopItem.Attribute {
	case types.StrengthAttribute:
		player.Strength += shopItem.Modifier
	case types.MovementAttribute:
		player.MoveCapacity += shopItem.Modifier
	case types.AttackVelocityAttribute:
		player.TotalShots += shopItem.Modifier
	default:
		return errors.New("Atributo não reconhecido")
	}
	return nil
}
