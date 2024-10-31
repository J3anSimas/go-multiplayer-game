package models

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAbc(t *testing.T) {
	assert := assert.New(t)
	games := make([]Room, 0)
	new_game, err := NewRoom(20, 20)
	assert.NoError(err, "Erro ao criar a sala")

	games = append(games, new_game)

	invite_code := new_game.GetInviteCode()
	room := GetRoomByInviteCode(&games, invite_code)
	_, err = room.JoinGame()
	assert.NoError(err, "Erro ao entrar na sala")

}
