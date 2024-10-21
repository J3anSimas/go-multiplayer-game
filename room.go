package main

import (
	"errors"

	"github.com/google/uuid"
)

type Status uint8

type Point struct {
	x, y int
}

var directions = []Point{
	{-1, 0},  // cima
	{1, 0},   // baixo
	{0, -1},  // esquerda
	{0, 1},   // direita
	{-1, -1}, // cima-esquerda (diagonal)
	{-1, 1},  // cima-direita (diagonal)
	{1, -1},  // baixo-esquerda (diagonal)
	{1, 1},   // baixo-direita (diagonal)
}

const (
	WorldDefaultWidth  int = 100
	WorldDefaultHeight int = 100

	PlayerStartingHealth         int = 10
	PlayerStartingMoveCapacity   int = 10
	PlayerStartingMovesRemaining int = 10
	PlayerStartingStrength       int = 10
	PlayerStartingTotalShots     int = 10
	PlayerStartingShotsRemaining int = 10
)
const (
	WaitingForConnection Status = iota
	Running
	Paused
	GameOver
)

type GameLogic struct {
	Room *Room
}

func (g *GameLogic) MovePlayer(playerId string, dx, dy int) ([]Point, error) {
	player := g.Room.FindPlayerById(playerId)
	if player == nil {
		return nil, errors.New("Jogador não encontrado")
	}
	return player.Move(dx, dy, g.Room)
}

type Room struct {
	Id          string
	Status      Status
	WorldWidth  int
	WorldHeight int
	Players     []Player
	Mobs        []Mob
}
type Player struct {
	Id             string `json:"id"`
	Ready          bool   `json:"ready"`
	IsHost         bool
	Position       Point
	Health         int
	MoveCapacity   int
	MovesRemaining int
	Strength       int
	TotalShots     int
	ShotsRemaining int
}
type Mob struct {
	Health   int
	Position Point
	Strength int
}

func isValid(x, y, nRows, nCols int, room Room) bool {
	if x < 0 || y < 0 || x > nRows || y > nCols {
		return false
	}
	for _, p := range room.Players {
		if p.Position.x == x && p.Position.y == y {
			return false
		}
	}
	for _, m := range room.Mobs {
		if m.Position.x == x && m.Position.y == y {
			return false
		}
	}
	return true
}

// FindShortestPath encontra o menor número de movimentos de start até end na matriz.
func FindShortestPath(room Room, startPoint, endPoint Point) (int, []Point) {
	nRows := room.WorldHeight
	nCols := room.WorldWidth

	queue := []Point{startPoint}
	visited := make([][]bool, nRows)
	for i := range visited {
		visited[i] = make([]bool, nCols)
	}
	visited[startPoint.x][startPoint.y] = true

	// Matriz para contar o número de passos e armazenar o caminho
	steps := make([][]int, nRows)
	for i := range steps {
		steps[i] = make([]int, nCols)
	}

	// Para reconstruir o caminho
	previous := make([][]Point, nRows)
	for i := range previous {
		previous[i] = make([]Point, nCols)
	}

	// BFS
	for len(queue) > 0 {
		currentPoint := queue[0]
		queue = queue[1:]

		if currentPoint.x == endPoint.x && currentPoint.y == endPoint.y {
			// Chegamos ao destino; reconstruir o caminho
			path := []Point{}
			for p := currentPoint; p != startPoint; p = previous[p.x][p.y] {
				path = append(path, p)
			}
			path = append(path, startPoint) // Adiciona o ponto de partida
			// Inverter o caminho para começar do início
			for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
				path[i], path[j] = path[j], path[i]
			}
			return steps[currentPoint.x][currentPoint.y], path
		}

		// Testar todas as 8 direções
		for _, direction := range directions {
			newX := currentPoint.x + direction.x
			newY := currentPoint.y + direction.y

			if isValid(newX, newY, nRows, nCols, room) && !visited[newX][newY] {
				visited[newX][newY] = true
				queue = append(queue, Point{newX, newY})
				steps[newX][newY] = steps[currentPoint.x][currentPoint.y] + 1
				previous[newX][newY] = currentPoint // Armazena o ponto anterior
			}
		}
	}

	return -1, nil // Se não houver caminho
}

func (p *Player) Move(posX, posY int, room *Room) ([]Point, error) {
	if posX < 0 || posX >= room.WorldWidth || posY < 0 || posY >= room.WorldHeight {
		return nil, errors.New("Movimento fora dos limites")
	}
	dest := Point{
		x: posX,
		y: posY,
	}

	distance, path := FindShortestPath(*room, p.Position, dest)
	if distance > p.MovesRemaining {
		return nil, errors.New("O jogador não possui movimentos suficientes")
	}

	p.Position = dest
	p.MovesRemaining -= distance
	return path, nil
}

func NewRoom(width, height int) (Room, error) {
	if width == 0 {
		width = WorldDefaultWidth
	}
	if height == 0 {
		height = WorldDefaultHeight
	}
	id := uuid.NewString()
	room := Room{
		Id:          id,
		WorldWidth:  width,
		WorldHeight: height,
		Status:      WaitingForConnection,
	}
	players := make([]Player, 1)
	players[0] = Player{
		Id:             uuid.NewString(),
		Ready:          false,
		IsHost:         true,
		Position:       Point{1, 1},
		Health:         PlayerStartingHealth,
		MoveCapacity:   PlayerStartingMoveCapacity,
		MovesRemaining: PlayerStartingMovesRemaining,
		Strength:       PlayerStartingStrength,
		TotalShots:     PlayerStartingTotalShots,
		ShotsRemaining: PlayerStartingShotsRemaining,
	}
	room.Players = players

	return room, nil
}
func (r *Room) FindPlayerById(playerId string) *Player {
	for i, p := range r.Players {
		if p.Id == playerId {
			return &r.Players[i]
		}
	}
	return nil
}
func (r Room) GetInviteCode() string {
	return r.Id[14:23]
}

func (r *Room) JoinGame() (Player, error) {
	player := Player{
		Id:             uuid.NewString(),
		Ready:          false,
		IsHost:         false,
		Position:       Point{r.WorldWidth, r.WorldHeight},
		Health:         PlayerStartingHealth,
		MoveCapacity:   PlayerStartingMoveCapacity,
		MovesRemaining: PlayerStartingMovesRemaining,
		Strength:       PlayerStartingStrength,
		TotalShots:     PlayerStartingTotalShots,
		ShotsRemaining: PlayerStartingShotsRemaining,
	}
	r.Players = append(r.Players, player)
	return player, nil
}

func (r *Room) StartGame() error {
	if len(r.Players) < 2 {
		return errors.New("São necessários 2 jogadores conectados")
	}
	for _, p := range r.Players {
		if !p.Ready {
			return errors.New("Nem todos os jogadores estão prontos")
		}
	}
	r.Status = Running
	return nil
}
func (r *Room) Update() {
}

func (p *Player) ToggleReady() {
	p.Ready = !p.Ready
}
