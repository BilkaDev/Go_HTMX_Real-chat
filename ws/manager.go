package ws

import (
	"context"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	components "github.com/bilkadev/Go_HTMX_Real-chat/view/components/sidebar"
	"github.com/gorilla/websocket"
	"github.com/labstack/echo/v4"
)

var (
	upgrader = websocket.Upgrader{}
)

type Manager struct {
	ClientList             []*Client
	ClientListEventChannel chan *ClientListEvent
}

func NewManager() *Manager {
	return &Manager{
		ClientList:             []*Client{},
		ClientListEventChannel: make(chan *ClientListEvent),
	}
}

func (m *Manager) HandleClientListEventChannel() {
	for {
		select {
		case ClientListEvent, ok := <-m.ClientListEventChannel:
			if !ok {
				return
			}
			switch ClientListEvent.EventType {
			case "ADD":
				for _, client := range m.ClientList {
					if client.ID == ClientListEvent.Client.ID {
						return
					}
				}
				m.ClientList = append(m.ClientList, ClientListEvent.Client)
				SendToUsersStatus(m.ClientList)
			case "REMOVE":
				newSlice := []*Client{}
				for _, client := range m.ClientList {
					if client.ID == ClientListEvent.Client.ID {
						continue
					}
					newSlice = append(newSlice, client)
				}
				SendToUsersStatus(newSlice)
				m.ClientList = newSlice
			}
		case <-context.Background().Done():
			return
		}
	}
}

func (m *Manager) Handler(c echo.Context) error {

	ws, err := upgrader.Upgrade(c.Response(), c.Request(), nil)
	if err != nil {
		return err
	}
	userID := c.Get(config.CurrentUserId.String()).(uint)
	userName := c.Get(config.CurrentUserName.String()).(string)
	nc := NewClient(ws, m, userID, userName)

	m.ClientListEventChannel <- &ClientListEvent{
		EventType: "ADD",
		Client:    nc,
	}

	go nc.WriteMessage(c)

	return nil
}
func SendUserStatus(c *Client, data string) {
	c.MessageChannel <- components.UsersOnline(data)
}

func SendToUsersStatus(cl []*Client) {
	var data string = ""
	for _, cl := range cl {
		data = data + "," + cl.UserName
	}
	for _, c := range cl {
		go SendUserStatus(c, data)
	}
}
