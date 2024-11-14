package models

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strconv"

	"github.com/J3anSimas/game_multiplayer_go/types"
)

type Player struct {
	Id             string
	Ready          bool
	IsHost         bool
	Position       types.Point
	Health         int
	Coins          int
	MoveCapacity   int
	MovesRemaining int
	Strength       int
	TotalShots     int
	ShotsRemaining int

	IsDead bool
}

func (p *Player) Move(posX, posY int, room *Room) ([]types.Point, error) {
	if posX < 0 || posX >= room.WorldWidth {
		return nil, errors.New("Pos X fora dos limites: " + strconv.Itoa(posX))
	}
	if posY < 0 || posY >= room.WorldHeight {
		return nil, errors.New("Pos Y fora dos limites: " + strconv.Itoa(posY))
	}
	dest := types.Point{
		X: posX,
		Y: posY,
	}

	distance, path := p.FindShortestPath(*room, dest)
	if distance == -1 {
		return nil, errors.New("Falha ao mover jogador")
	}
	if distance > p.MovesRemaining {
		return nil, errors.New("o jogador não possui movimentos suficientes")
	}

	p.Position = dest
	p.MovesRemaining -= distance
	return path, nil
}

func (p *Player) ToggleReady() {
	p.Ready = !p.Ready
}

var directions = []types.Point{
	{X: -1, Y: 0},  // cima
	{X: 1, Y: 0},   // baixo
	{X: 0, Y: -1},  // esquerda
	{X: 0, Y: 1},   // direita
	{X: -1, Y: -1}, // cima-esquerda (diagonal)
	{X: -1, Y: 1},  // cima-direita (diagonal)
	{X: 1, Y: -1},  // baixo-esquerda (diagonal)
	{X: 1, Y: 1},   // baixo-direita (diagonal)
}

func isValid(x, y, nRows, nCols int, room Room) bool {
	if x < 0 || y < 0 || x >= nRows || y >= nCols {
		return false
	}
	for _, p := range room.Players {
		if p.Position.X == x && p.Position.Y == y {
			return false
		}
	}
	for _, m := range room.Mobs {
		if m.Position.X == x && m.Position.Y == y {
			return false
		}
	}
	return true
}

// FindShortestPath encontra o menor número de movimentos de start até end na matriz.
func (p Player) FindShortestPath(room Room, endPoint types.Point) (int, []types.Point) {
	nRows := room.WorldHeight
	nCols := room.WorldWidth

	queue := []types.Point{p.Position}
	visited := make([][]bool, nRows)
	for i := range visited {
		visited[i] = make([]bool, nCols)
	}
	visited[p.Position.X][p.Position.Y] = true

	// Matriz para contar o número de passos e armazenar o caminho
	steps := make([][]int, nRows)
	for i := range steps {
		steps[i] = make([]int, nCols)
	}

	// Para reconstruir o caminho
	previous := make([][]types.Point, nRows)
	for i := range previous {
		previous[i] = make([]types.Point, nCols)
	}

	// BFS
	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		if currentPoint.X == endPoint.X && currentPoint.Y == endPoint.Y {
			// Chegamos ao destino; reconstruir o caminho
			path := []types.Point{}
			for pth := currentPoint; pth != p.Position; pth = previous[pth.X][pth.Y] {
				path = append(path, pth)
			}
			path = append(path, p.Position) // Adiciona o ponto de partida
			// Inverter o caminho para começar do início
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return steps[currentPoint.X][currentPoint.Y], path[1:]
		}

		// Testar todas as 8 direções
		for _, direction := range directions {
			newX := currentPoint.X + direction.X
			newY := currentPoint.Y + direction.Y

			if isValid(newX, newY, nRows, nCols, room) && !visited[newX][newY] {
				visited[newX][newY] = true
				queue = append(queue, types.Point{X: newX, Y: newY})
				steps[newX][newY] = steps[currentPoint.X][currentPoint.Y] + 1
				previous[newX][newY] = currentPoint // Armazena o ponto anterior
			}
		}
	}

	return -1, nil // Se não houver caminho
}
func (p *Player) ResetAttributes() {
	p.ShotsRemaining = p.TotalShots
	p.MovesRemaining = p.MoveCapacity
}
