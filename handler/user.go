package handler

import (
	"fmt"
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
	q := c.QueryParam("search")
	fmt.Println("query:", q)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Can't parse to uint")
	}

	users, err := h.store.FindAllWithoutSender(currentUserId, q)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}

	return render(c, components.Conversations(*users))
}
