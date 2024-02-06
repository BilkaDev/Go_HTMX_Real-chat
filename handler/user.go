package handler

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/layout"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{Email: "email@exmaple.pl"}
	return render(c, user.Show(u), layout.Base())
}
