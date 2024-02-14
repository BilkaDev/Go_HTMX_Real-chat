package handler

import (
	"fmt"
	"net/http"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/bilkadev/Go_HTMX_Real-chat/middleware"
	"github.com/bilkadev/Go_HTMX_Real-chat/store"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/home"
	"github.com/labstack/echo/v4"
)

type HomeHandler struct {
	StoreUser *store.UserStore
}

func HomeRouter(e *echo.Group, prefix string, storage *store.SqlStore) {
	h := &HomeHandler{
		StoreUser: store.NewUserStore(storage),
	}

	g := e.Group(prefix)
	g.Use(middleware.RequireAuth)

	g.GET("", h.handleHomeShow)
}

func (h *HomeHandler) handleHomeShow(c echo.Context) error {
	cuId := c.Get(config.CurrentUserId.String()).(uint)
	u, err := h.StoreUser.FindAllWithoutSender(cuId, "")
	if err != nil {
		fmt.Println(err.Error())
		c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER Something went wrong!")
	}
	return render(c, home.Show(*u, true))
}
