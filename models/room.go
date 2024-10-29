package models

import (
	"errors"

	"github.com/J3anSimas/game_multiplayer_go/types"
	"github.com/J3anSimas/game_multiplayer_go/utils"
	"github.com/google/uuid"
)

type Room struct {
	Id          string
	Status      types.Status
	WorldWidth  int
	WorldHeight int
	Players     []*Player
	Mobs        []*Mob
	ShopItems   []ShopItem
	Turn        types.Turn
}

func NewRoom(width, height int) (Room, error) {
	if width == 0 {
		width = utils.WorldDefaultWidth
	}
	if height == 0 {
		height = utils.WorldDefaultHeight
	}
	id := uuid.NewString()
	shopItems := make([]ShopItem, 3)
	shopItems[0] = ShopItem{
		Title:       "Botas de Movimento",
		Description: "Aumenta o número de movimentos por turno em 2",
		Cost:        10,
		Attribute:   types.MovementAttribute,
		Modifier:    2,
	}
	shopItems[1] = ShopItem{
		Title:       "Lâmina do Infinito",
		Description: "Aumenta a força em 5",
		Cost:        10,
		Attribute:   types.StrengthAttribute,
		Modifier:    5,
	}
	shopItems[2] = ShopItem{
		Title:       "Dançarina Fantasma",
		Description: "Aumenta o número de ataques por turno em 1",
		Cost:        17,
		Attribute:   types.AttackVelocityAttribute,
		Modifier:    1,
	}
	room := Room{
		Id:          id,
		WorldWidth:  width,
		WorldHeight: height,
		Status:      types.WaitingForConnection,
		Turn:        types.HostTurn,
	}
	players := make([]*Player, 1)
	players[0] = &Player{
		Id:             uuid.NewString(),
		Ready:          false,
		IsHost:         true,
		Position:       types.Point{X: 0, Y: 0},
		Health:         utils.PlayerStartingHealth,
		MoveCapacity:   utils.PlayerStartingMoveCapacity,
		MovesRemaining: utils.PlayerStartingMovesRemaining,
		Strength:       utils.PlayerStartingStrength,
		TotalShots:     utils.PlayerStartingTotalShots,
		ShotsRemaining: utils.PlayerStartingShotsRemaining,
		IsDead:         false,
	}
	room.Players = players
	room.GenerateMobs()
	return room, nil
}

func (r *Room) FindPlayerById(playerId string) *Player {
	for i, p := range r.Players {
		if p.Id == playerId {
			return r.Players[i]
		}
	}
	return nil
}

func (r *Room) JoinGame() (Player, error) {
	player := Player{
		Id:             uuid.NewString(),
		Ready:          false,
		IsHost:         false,
		Position:       types.Point{X: r.WorldWidth - 1, Y: r.WorldHeight - 1},
		Health:         utils.PlayerStartingHealth,
		MoveCapacity:   utils.PlayerStartingMoveCapacity,
		MovesRemaining: utils.PlayerStartingMovesRemaining,
		Strength:       utils.PlayerStartingStrength,
		TotalShots:     utils.PlayerStartingTotalShots,
		ShotsRemaining: utils.PlayerStartingShotsRemaining,
		IsDead:         false,
	}
	r.Players = append(r.Players, &player)
	return player, nil
}

func (r *Room) StartGame() error {
	if len(r.Players) < 2 {
		return errors.New("são necessários 2 jogadores conectados")
	}
	for _, p := range r.Players {
		if !p.Ready {
			return errors.New("nem todos os jogadores estão prontos")
		}
	}
	r.Status = types.Running
	r.Turn = types.HostTurn
	return nil
}
func (r *Room) GenerateMobs() {
	mobs := make([]*Mob, 4)
	posX := int(float64(r.WorldWidth) * 0.3)
	posY := int(float64(r.WorldHeight) * 0.1)
	mobs[0] = NewMob(15, types.Point{X: posX, Y: posY}, 5, 5)

	posX = int(float64(r.WorldWidth) * 0.7)
	posY = int(float64(r.WorldHeight) * 0.9)
	mobs[1] = NewMob(15, types.Point{X: posX, Y: posY}, 5, 5)

	posX = int(float64(r.WorldWidth) * 0.5)
	posY = int(float64(r.WorldHeight) * 0.0)
	mobs[2] = NewMob(15, types.Point{X: posX, Y: posY}, 5, 5)

	posX = int(float64(r.WorldWidth) * 0.5)
	posY = int(float64(r.WorldHeight)*1.0) - 1
	mobs[3] = NewMob(15, types.Point{X: posX, Y: posY}, 5, 5)

	r.Mobs = mobs
}
