package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

func main() {
	rooms := make([]Room, 0)
	http.HandleFunc("/room", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if r.Method == http.MethodPost {
			player := Player{}
			err := json.NewDecoder(r.Body).Decode(&player)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				resp := ErrorResponse{
					Error: "Failed to read user data",
				}
				err = json.NewEncoder(w).Encode(resp)
				if err != nil {
					http.Error(w, "Failed to generate json message", http.StatusInternalServerError)

				}
				return
			}
			room, err := NewRoom(player)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				resp := ErrorResponse{
					Error: "Failed to create room",
				}
				err = json.NewEncoder(w).Encode(resp)
				if err != nil {
					http.Error(w, "Failed to generate json message", http.StatusInternalServerError)
				}
				return
			}
			rooms = append(rooms, room)
			fmt.Fprintf(w, "{\"gameId\": \"%s\",\n\"gameCode\": \"%s\"}", room.Id, room.GetGameCode())
			return

		} else {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
	})
	fmt.Println("Server listening on port :8080")
	http.ListenAndServe(":8080", nil)
}

type Player struct {
	Id string `json:"id"`
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
func (r Room) GetGameCode() string {
	return r.Id[14:23]
}
