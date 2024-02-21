package handler

import (
	"net/http"
	"time"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/bilkadev/Go_HTMX_Real-chat/middleware"
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"github.com/bilkadev/Go_HTMX_Real-chat/pkg"
	"github.com/bilkadev/Go_HTMX_Real-chat/store"
	components "github.com/bilkadev/Go_HTMX_Real-chat/view/components/messages"
	"github.com/bilkadev/Go_HTMX_Real-chat/ws"
	"github.com/labstack/echo/v4"
)

type MessageHandler struct {
	StoreMessage      *store.MessageStore
	StoreUser         *store.UserStore
	StoreConversation *store.ConversationStore
	Socket            *ws.Manager
}

func MessageRouter(e *echo.Group, prefix string, storage *store.SqlStore, socket *ws.Manager) {
	mh := &MessageHandler{
		StoreMessage:      store.NewMessageStore(storage),
		StoreUser:         store.NewUserStore(storage),
		StoreConversation: store.NewConversationStore(storage),
		Socket:            socket,
	}

	g := e.Group(prefix)
	g.Use(middleware.RequireAuth)

	g.GET("/:id", mh.handleMessagesShow)
	g.POST("/send/:id", mh.handleMessageSend)
}

func (h *MessageHandler) handleMessageSend(c echo.Context) error {
	var m = model.Message{}
	err := pkg.FormValidate(c, &m)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	receiver, err := h.StoreUser.FindOne(m.ReceiverID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	currentUserId := c.Get(config.CurrentUserId.String()).(uint)
	m.SenderID = currentUserId
	sender, err := h.StoreUser.FindOne(currentUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	conversation := model.Conversation{}
	h.StoreConversation.FindOneOrCreate(&conversation, m)

	// @Todo socket IO func will go here
	for _, client := range h.Socket.ClientList {
		if client.UserID == receiver.ID {
			m.CreatedAt = time.Now()
			client.MessageChannel <- components.Message(false, *sender, *receiver, m, true)
		}
	}

	// push to conversation
	conversation.Messages = append(conversation.Messages, m)
	if _, err = h.StoreConversation.Update(&conversation); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	// res 201 return new message
	return render(c, components.Message(true, *sender, *receiver, m, false))
}

func (h *MessageHandler) handleMessagesShow(c echo.Context) error {
	m := model.Message{}
	err := (&echo.DefaultBinder{}).BindPathParams(c, &m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	currentUserId := c.Get(config.CurrentUserId.String()).(uint)

	receiver, err := h.StoreUser.FindOne(m.ReceiverID)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	currentUser, err := h.StoreUser.FindOne(currentUserId)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	conversation := model.Conversation{}

	if _, err := h.StoreConversation.FindOne(&conversation, []uint{
		currentUserId,
		m.ReceiverID,
	}); err != nil {
		return render(c, components.Messages(currentUser, receiver, &[]model.Message{}))
	}

	messages, err := h.StoreMessage.FindAll(conversation.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return render(c, components.Messages(currentUser, receiver, messages))
}
