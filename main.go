package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"github.com/J3anSimas/game_multiplayer_go/models"
	"github.com/J3anSimas/game_multiplayer_go/types"
)

var (
	PORT = 8080
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	room, err := models.NewRoom(10, 10)
	if err != nil {
		panic(err)
	}
	_ = models.GameLogic{
		Room: &room,
	}
	room.JoinGame()
	room.Players[0].ToggleReady()
	room.Players[1].ToggleReady()
	err = room.StartGame()
	if err != nil {
		panic(err)
	}
	for {
		var cmd string
		var currentPlayer string
		switch room.Turn {
		case types.HostTurn:
			currentPlayer = "Player 1"
		default:
			currentPlayer = "Player 2"
		}
		fmt.Printf("%s, entre com o comando: ", currentPlayer)
		if scanner.Scan() {
			cmd = scanner.Text()
		}
		if err := scanner.Err(); err != nil {
			fmt.Println("Algo deu errado", err.Error())
			continue
		}
		fmt.Printf("%s\n", cmd)
	}
	// PrintWorld(room)
	// gl.MovePlayer(gl.Room.Players[0], 10, 10)

}
func PrintWorld(r models.Room) {
	w := make([][]string, r.WorldWidth)
	for i := range w {
		w[i] = make([]string, r.WorldHeight)
	}
	for i, p := range r.Players {
		w[p.Position.X][p.Position.Y] = "P" + strconv.Itoa(i+1)
	}
	for i, m := range r.Mobs {
		w[m.Position.X][m.Position.Y] = "M" + strconv.Itoa(i+1)
	}

	for _, i := range w {
		for _, j := range i {
			fmt.Printf(j + " ")
		}
		fmt.Println()
	}

}
