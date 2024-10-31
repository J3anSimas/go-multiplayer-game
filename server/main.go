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
	conn *websocket.Conn
	room string
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
			if h.rooms[client.room] == nil {
				h.rooms[client.room] = make(map[*websocket.Conn]bool)
			}
			h.rooms[client.room][client.conn] = true

		case client := <-h.unregister:
			if clients, ok := h.rooms[client.room]; ok {
				if _, ok := clients[client.conn]; ok {
					delete(clients, client.conn)
					if len(clients) == 0 {
						delete(h.rooms, client.room)
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
		return c.String(200, "Join")
	})
	e.GET("/ws/:room", func(c echo.Context) error {
		room := c.Param("room")
		conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
		if err != nil {
			// panic(err)
			log.Printf("%s, error while Upgrading websocket connection\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		client := &Client{conn: conn, room: room}
		hub.register <- client
		defer func() {
			hub.unregister <- client
			conn.Close()
		}()

		for {
			// Read message from client
			messageType, p, err := conn.ReadMessage()
			if err != nil {
				// panic(err)
				log.Printf("%s, error while reading message\n", err.Error())
				return c.JSON(http.StatusInternalServerError, map[string]string{
					"error": err.Error(),
				})
			}

			type Message struct {
				CardId string `json:"cardId"`
				Status string `json:"status"`
			}
			msg := Message{}
			json.Unmarshal(p, &msg)
			fmt.Println("Messagem recebida", msg.Status, msg.CardId)
			if clients, ok := hub.rooms[client.room]; ok {
				for client := range clients {
					if client != conn {
						err = client.WriteMessage(messageType, p)
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
			err = conn.WriteMessage(messageType, p)
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
