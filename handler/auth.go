package handler

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/auth"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/layout"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/user"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct{}

func AuthRouter(e *echo.Echo, prefix string) {
	ah := AuthHandler{}
	g := e.Group(prefix)
	g.GET("", ah.HanndleAuthShow)
	g.POST("/login", ah.HandleAuthLogin)
	g.POST("/SignUp", ah.HandleAuthSignUp)
	g.POST("/logout", ah.HandleAuthLogout)
}

func (h AuthHandler) HanndleAuthShow(c echo.Context) error {
	return render(c, auth.Show(), layout.Base())
}

func (h AuthHandler) HandleAuthLogin(c echo.Context) error {
	u := model.User{
		Email: "test",
	}
	return render(c, user.Show(u), layout.Base())
}

func (h AuthHandler) HandleAuthLogout(c echo.Context) error {
	u := model.User{
		Email: "test",
	}
	return render(c, user.Show(u), layout.Base())
}

func (h AuthHandler) HandleAuthSignUp(c echo.Context) error {
	u := model.User{
		Email: "test",
	}
	return render(c, user.Show(u), layout.Base())
}
