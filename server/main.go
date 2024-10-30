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
	game_logics := make([]*models.GameLogic, 0)
	e := echo.New()
	t := &Template{
		templates: template.Must(template.ParseGlob("public/views/*.html")),
	}
	e.Renderer = t

	e.GET("/", func(c echo.Context) error {
		return c.Render(http.StatusOK, "index.html", "Teste")
	})
	e.POST("/room", func(c echo.Context) error {
		fmt.Printf("Teste")
		new_room, err := models.NewRoom(100, 100)
		if err != nil {
			c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		game_logic := models.GameLogic{
			Room: &new_room,
		}
		invite_code := game_logic.Room.GetInviteCode()
		game_logics = append(game_logics, &game_logic)

		conn, err := upgrader.Upgrade(c.Response().Writer, c.Request(), nil)
		if err != nil {
			// panic(err)
			log.Printf("%s, error while Upgrading websocket connection\n", err.Error())
			return c.JSON(http.StatusInternalServerError, map[string]string{
				"error": err.Error(),
			})
		}
		client := &Client{conn: conn, room: game_logic.Room.Id}
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
				break
			}

			type Message struct {
				CardId string `json:"cardId"`
				Status string `json:"status"`
			}
			msg := Message{}
			json.Unmarshal(p, &msg)
			fmt.Println("Messagem recebida", msg.Status, msg.CardId)
			if clients, ok := hub.rooms[room]; ok {
				for client := range clients {
					if client != conn {
						err = client.WriteMessage(messageType, p)
						if err != nil {
							// panic(err)
							log.Printf("%s, error while writing message\n", err.Error())
							return c.JSON(http.StatusInternalServerError, map[string]string{
								"error": err.Error(),
							})
							break
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
				break
			}
		}
		return c.JSON(http.StatusCreated,
			map[string]string{"invite_code": invite_code})
	})
	e.Logger.Fatal(e.Start(fmt.Sprintf(":%d", PORT)))
}
