package main

import (
	"fmt"

	"github.com/J3anSimas/game_multiplayer_go/models"
	"github.com/J3anSimas/game_multiplayer_go/types"
)

var (
	PORT = 8080
)

func main() {
	room, _ := models.NewRoom(5, 5)
	mobs := make([]*models.Mob, 1)
	mobs[0] = &models.Mob{
		Health:   100,
		Position: types.Point{X: 3, Y: 3},
		Strength: 10,
	}
	room.Mobs = mobs
	room.Players[0].Position = types.Point{X: 0, Y: 0}
	steps, err := room.Players[0].Move(4, 3, &room)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Moves remaining: %d\nCurrent Position: %d, %d\nSteps: %v#\n",
		room.Players[0].MovesRemaining, room.Players[0].Position.X,
		room.Players[0].Position.Y, steps)
	room.Players[0].MovesRemaining = 6
	steps, err = room.Players[0].Move(4, 3, &room)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Printf("Moves remaining: %d\nCurrent Position: %d, %d\nSteps: %v#\n",
		room.Players[0].MovesRemaining, room.Players[0].Position.X,
		room.Players[0].Position.Y, steps)
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
