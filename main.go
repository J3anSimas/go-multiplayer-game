package main

import (
	"fmt"
)

var (
	PORT = 8080
)

func main() {
	// rooms := make([]Room, 0)
	// players := make([]Player, 0)
	// e := echo.New()
	// e.POST("/get-credentials", func(c echo.Context) error {
	// 	var player Player
	// 	if err := c.Bind(&player); err != nil {
	// 		return echo.NewHTTPError(http.StatusBadRequest, "Input inválido")
	// 	}
	// 	player = NewPlayer(player.Id, player.Username)
	// 	players = append(players, player)
	// 	return c.JSON(http.StatusOK, player)
	//
	// })
	// e.POST("/room", func(c echo.Context) error {
	// 	c.Response().Header().Set("Content-Type", "application/json")
	// 	player := Player{}
	// 	if err := c.Bind(&player); err != nil {
	// 		return echo.NewHTTPError(http.StatusBadRequest, "Failed to read user data")
	//
	// 	}
	// 	room, err := NewRoom(player)
	// 	if err != nil {
	// 		return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create room")
	// 	}
	// 	rooms = append(rooms, room)
	//
	// 	return c.JSON(http.StatusCreated, map[string]string{"roomId": room.Id, "roomCode": room.GetInviteCode()})
	//
	// })
	fmt.Printf("Server listening on port :%d\n", PORT)
	// e.Start(fmt.Sprintf(":%d", PORT))
}
