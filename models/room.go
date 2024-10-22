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
}

func NewRoom(width, height int) (Room, error) {
	if width == 0 {
		width = utils.WorldDefaultWidth
	}
	if height == 0 {
		height = utils.WorldDefaultHeight
	}
	id := uuid.NewString()
	room := Room{
		Id:          id,
		WorldWidth:  width,
		WorldHeight: height,
		Status:      types.WaitingForConnection,
	}
	players := make([]*Player, 1)
	players[0] = &Player{
		Id:             uuid.NewString(),
		Ready:          false,
		IsHost:         true,
		Position:       types.Point{X: 1, Y: 1},
		Health:         utils.PlayerStartingHealth,
		MoveCapacity:   utils.PlayerStartingMoveCapacity,
		MovesRemaining: utils.PlayerStartingMovesRemaining,
		Strength:       utils.PlayerStartingStrength,
		TotalShots:     utils.PlayerStartingTotalShots,
		ShotsRemaining: utils.PlayerStartingShotsRemaining,
	}
	room.Players = players

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
		Position:       types.Point{X: r.WorldWidth, Y: r.WorldHeight},
		Health:         utils.PlayerStartingHealth,
		MoveCapacity:   utils.PlayerStartingMoveCapacity,
		MovesRemaining: utils.PlayerStartingMovesRemaining,
		Strength:       utils.PlayerStartingStrength,
		TotalShots:     utils.PlayerStartingTotalShots,
		ShotsRemaining: utils.PlayerStartingShotsRemaining,
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
	return nil
}
