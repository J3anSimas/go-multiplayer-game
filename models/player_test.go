package models

import (
	"strconv"
	"testing"

	"github.com/J3anSimas/game_multiplayer_go/types"
	"github.com/stretchr/testify/assert"
)

// func TestPlayerMove(t *testing.T) {
// 	room, err := NewRoom(5, 5)
// 	assert.NoError(t, err, "Erro ao criar a sala")
//
// 	player := room.Players[0]
// 	player.Position = types.Point{X: 1, Y: 1}
// 	player.MovesRemaining = 5
//
// 	tests := []struct {
// 		name      string
// 		moveX     int
// 		moveY     int
// 		expectPos types.Point
// 		expectErr bool
// 	}{
// 		{"Movimento Válido", 2, 2, types.Point{X: 2, Y: 2}, false},
// 		{"Movimento Fora dos Limites", 6, 6, types.Point{}, true},
// 		{"Movimento Insuficiente", 3, 3, types.Point{}, true}, // Movendo 3 posições
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			// Faz a movimentação
// 			_, err := player.Move(tt.moveX, tt.moveY, &room)
//
// 			if tt.expectErr {
// 				assert.Error(t, err, "Esperado erro, mas não ocorreu.")
// 			} else {
// 				assert.NoError(t, err, "Erro inesperado: %v", err)
// 				assert.Equal(t, tt.expectPos, player.Position, "Posição esperada não corresponde.")
// 			}
// 		})
// 	}
// }

func TestPlayerMovementNew(t *testing.T) {
	assert := assert.New(t)
	room, _ := NewRoom(5, 5)
	mobs := make([]*Mob, 1)
	mobs[0] = &Mob{
		Health:   100,
		Position: types.Point{X: 3, Y: 3},
		Strength: 10,
	}
	room.Mobs = mobs
	room.Players[0].Position = types.Point{X: 0, Y: 0}
	room.Players[0].MovesRemaining = 3
	_, err := room.Players[0].Move(4, 3, &room)
	assert.NotNil(err, "Esperado que o room.Players[0] não tenha movimento suficiente")
	room.Players[0].MovesRemaining = 4
	_, err = room.Players[0].Move(4, 3, &room)
	assert.Nil(err, "Esperado nenhum erro, mas algo aconteceu: "+err.Error())
	movesRemainingString := strconv.Itoa(room.Players[0].MovesRemaining)
	assert.Equal(0, room.Players[0].MovesRemaining, "Número de movimentos errado: "+movesRemainingString)

}
