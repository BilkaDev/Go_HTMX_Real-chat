package handler

import (
	"net/http"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/bilkadev/Go_HTMX_Real-chat/middleware"
	"github.com/bilkadev/Go_HTMX_Real-chat/store"
	components "github.com/bilkadev/Go_HTMX_Real-chat/view/components/sidebar"
	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	store *store.UserStore
}

func UserRouter(e *echo.Group, prefix string, storage *store.SqlStore) {
	uh := &UserHandler{
		store: store.NewUserStore(storage),
	}
	g := e.Group(prefix)
	g.Use(middleware.RequireAuth)
	g.GET("", uh.HandleUsersShow)
}

func (h UserHandler) HandleUsersShow(c echo.Context) error {
	currentUserId, ok := c.Get(config.CurrentUserId.String()).(uint)
	if !ok {
		return c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER can't parse to uint")
	}

	users, err := h.store.FindAllWithoutSender(currentUserId)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER "+err.Error())
	}

	return render(c, components.Conversations(*users))
}
