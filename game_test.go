package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlayerMovement(t *testing.T) {
	room, err := NewRoom()
	if err != nil {
		assert.Nil(t, err, "Failed to create room")
	}
	room.Players[0].Move(2, 3, &room)

}
