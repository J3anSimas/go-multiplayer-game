package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
	"net/http"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

var (
	PORT = 8080
)

func main() {
	rooms := make([]Room, 0)
	players := make([]Player, 0)
	e := echo.New()
	e.POST("/get-credentials", func(c echo.Context) error {
		var player Player
		if err := c.Bind(&player); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Input inválido")
		}
		player = NewPlayer(player.Id, player.Username)
		players = append(players, player)
		return c.JSON(http.StatusOK, player)

	})
	e.POST("/room", func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json")
		player := Player{}
		if err := c.Bind(&player); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, "Failed to read user data")

		}
		room, err := NewRoom(player)
		if err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, "Failed to create room")
		}
		rooms = append(rooms, room)

		return c.JSON(http.StatusCreated, map[string]string{"roomId": room.Id, "roomCode": room.GetInviteCode()})

	})
	fmt.Printf("Server listening on port :%d\n", PORT)
	e.Start(fmt.Sprintf(":%d", PORT))
}

type Player struct {
	Id       string `json:"id"`
	Username string `json:"username"`
}

func NewPlayer(id string, username string) Player {
	if id == "" {
		id = uuid.NewString()
	}
	return Player{
		id,
		username,
	}
}

type Room struct {
	Id      string
	Players []Player
}

func NewRoom(player Player) (Room, error) {
	id := uuid.NewString()
	players := make([]Player, 1)
	players[0] = player
	return Room{
		Id:      id,
		Players: players,
	}, nil

}
func (r Room) GetInviteCode() string {
	return r.Id[14:23]
}
