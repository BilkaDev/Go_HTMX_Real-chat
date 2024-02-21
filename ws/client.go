package ws

import (
	"bytes"
	"context"
	"time"

	"github.com/a-h/templ"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

type Client struct {
	ID             string
	UserID         uint
	Conn           *websocket.Conn
	Manager        *Manager
	MessageChannel chan templ.Component
	UserName       string
}

type ClientListEvent struct {
	EventType string
	Client    *Client
}

var tickerInterval = time.Second + 5

func NewClient(conn *websocket.Conn, manager *Manager, userId uint, userName string) *Client {
	return &Client{
		Conn:           conn,
		Manager:        manager,
		MessageChannel: make(chan templ.Component),
		ID:             uuid.New().String(),
		UserID:         userId,
		UserName:       userName,
	}
}

func (c *Client) WriteMessage(ctx echo.Context) {
	defer func() {
		c.Conn.Close()
		c.Manager.ClientListEventChannel <- &ClientListEvent{
			Client:    c,
			EventType: "REMOVE",
		}
	}()

	ticker := time.NewTicker(tickerInterval)

	for {
		select {
		case component, ok := <-c.MessageChannel:
			if !ok {
				return
			}
			buffer := &bytes.Buffer{}
			component.Render(context.Background(), buffer)
			err := c.Conn.WriteMessage(websocket.TextMessage, buffer.Bytes())
			if err != nil {
				ctx.Logger().Error(err)
				return
			}
		case <-context.Background().Done():
			return
		case <-ticker.C:
			if err := c.Conn.WriteMessage(websocket.PingMessage, []byte("")); err != nil {
				ctx.Logger().Error(err)
				return
			}
		}
	}
}
