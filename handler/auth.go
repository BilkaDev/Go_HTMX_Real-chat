package handler

import (
	"net/http"

	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"github.com/bilkadev/Go_HTMX_Real-chat/store"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/auth"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/layout"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/user"
	"github.com/labstack/echo/v4"
)

type AuthHandler struct {
	store *store.UserStore
}

func AuthRouter(e *echo.Echo, prefix string, storage *store.SqlStore) {
	ah := &AuthHandler{
		store: store.NewUserStore(storage),
	}
	g := e.Group(prefix)
	g.GET("", ah.HanndleAuthShow)
	g.POST("/login", ah.HandleAuthLogin)
	g.POST("/signup", ah.HandleAuthSignUp)
	g.POST("/logout", ah.HandleAuthLogout)
}

func (h *AuthHandler) HanndleAuthShow(c echo.Context) error {

	_ = model.User{FullName: "adam", Email: "tests@wdp.pl"}
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
	u := model.User{}
	err := c.Bind(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST")
	}
	// @Todo add validation
	_, ok := h.store.FindOneByEmail(u.Email)
	if ok == nil {
		return c.String(http.StatusConflict, "ERR_CONFLICT user with given email already exists")
	}
	// @Todo add passwod hash
	_, err = h.store.Create(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST"+err.Error())
	}
	// @Todo generate JWT

	return render(c, user.Show(u), layout.Base())
}
