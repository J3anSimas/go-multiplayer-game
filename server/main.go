package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"

	"github.com/J3anSimas/game_multiplayer_go/models"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	PORT = 8080
)
var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Template struct {
	templates *template.Template
}
type Client struct {
	conn     *websocket.Conn
	roomId   string
	playerId string
}
type Hub struct {
	rooms      map[string]map[*websocket.Conn]bool
	register   chan *Client
	unregister chan *Client
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			if h.rooms[client.roomId] == nil {
				h.rooms[client.roomId] = make(map[*websocket.Conn]bool)
			}
			h.rooms[client.roomId][client.conn] = true

		case client := <-h.unregister:
			if clients, ok := h.rooms[client.roomId]; ok {
				if _, ok := clients[client.conn]; ok {
					delete(clients, client.conn)
					if len(clients) == 0 {
						delete(h.rooms, client.roomId)
					}
				}
			}
		}
	}
}

func (t *Template) Render(w io.Writer, name string, data interface{}, c echo.Context) error {
	return t.templates.ExecuteTemplate(w, name, data)
}
func main() {
	hub := &Hub{
		rooms:      make(map[string]map[*websocket.Conn]bool),
		register:   make(chan *Client),
		unregister: make(chan *Client),
	}
	go hub.run()
	games := make([]models.Room, 0)
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "Teste")
	})
	e.POST("/room", func(c echo.Context) error {
		new_game, err := models.NewRoom(20, 20)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		invite_code := new_game.GetInviteCode()
		games = append(games, new_game)
		type Response struct {
			Game       models.Room `json:"game"`
			InviteCode string      `json:"invite_code"`
		}
		response := Response{
			Game:       new_game,
			InviteCode: invite_code,
		}
		return c.JSON(http.StatusCreated, response)

	})
	e.POST("/room/:invite_code/join", func(c echo.Context) error {
		invite_code := c.Param("invite_code")
		room := models.GetRoomByInviteCode(&games, invite_code)
		if room == nil {
			return c.JSON(http.StatusNotFound, map[string]string{
				"error": "Sala não encontrada",
			})
		}
		_, err := room.JoinGame()
		if err != nil {
			return c.JSON(http.StatusForbidden, map[string]string{
				"error": err.Error(),
			})
		}
		return c.JSON(http.StatusCreated, room)
	})
	e.GET("/ws/:room/:player", func(c echo.Context) error {
		room_id := c.Param("room")
		playerId := c.Param("player")
		conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
		if err != nil {
			// panic(err)
			log.Printf("%s, error while Upgrading websocket connection\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		client := &Client{conn: conn, roomId: room_id, playerId: playerId}
		hub.register <- client
		if clients, ok := hub.rooms[client.roomId]; ok {
			game_state := models.GetRoomById(games, room_id)
			json_data, err := json.Marshal(game_state)
			if err != nil {
				log.Printf("%s, Erro ao serializar para json\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
			for client := range clients {
				err = client.WriteMessage(websocket.TextMessage, json_data)
				if err != nil {
					// panic(err)
					log.Printf("%s, error while writing message\n", err.Error())
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": err.Error(),
					})
				}
			}
		}
		defer func() {
			hub.unregister <- client
			conn.Close()
		}()

		for {
			// Read message from client
			_, p, err := conn.ReadMessage()
			if err != nil {
				// panic(err)
				log.Printf("%s, error while reading message\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
			type Message struct {
				Cmd    string   `json:"cmd"`
				Params []string `json:"params"`
			}
			msg := Message{}
			err = json.Unmarshal(p, &msg)
			if err != nil {
				log.Printf("%s, Erro ao deserializar o json\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
			switch msg.Cmd {
			case "set_ready":
				room := models.GetRoomById(games, room_id)
				if room == nil {
					log.Printf("Jogo não encontrado\n")
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": "Jogo não encontrado",
					})
				}
				player := room.FindPlayerById(playerId)
				fmt.Println(playerId)
				if player == nil {
					log.Printf("Jogador não encontrado\n")
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": "Jogador não encontrado",
					})
				}
				err = room.TogglePlayerReady(player)
				if err != nil {
					log.Printf("Jogo não encontrado\n")
					return c.JSON(http.StatusInternalServerError, map[string]string{
						"error": "Jogo não encontrado",
					})
				}
				fmt.Println("Jogador Pronto", player.Ready)
			default:
				log.Printf("Comando não reconhecido\n")
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": "Comando não reconhecido",
				})

			}
			game_state := models.GetRoomById(games, room_id)
			json_data, err := json.Marshal(game_state)
			if err != nil {
				log.Printf("%s, Erro ao serializar para json\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
			fmt.Println("Messagem recebida", msg.Cmd, msg.Params)
			if clients, ok := hub.rooms[client.roomId]; ok {
				for client := range clients {
					if client != conn {
						err = client.WriteMessage(websocket.TextMessage, json_data)
						if err != nil {
							// panic(err)
							log.Printf("%s, error while writing message\n", err.Error())
							return c.JSON(http.StatusInternalServerError, map[string]string{
								"error": err.Error(),
							})
						}
					}
				}
			}
			// Echo message back to client
			err = conn.WriteMessage(websocket.TextMessage, json_data)
			if err != nil {
				// panic(err)
				log.Printf("%s, error while writing message\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}
		}
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
