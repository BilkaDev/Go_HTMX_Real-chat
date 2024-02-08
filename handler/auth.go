package handler

import (
	"fmt"
	"net/http"

	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"github.com/bilkadev/Go_HTMX_Real-chat/pkg"
	"github.com/bilkadev/Go_HTMX_Real-chat/pkg/security"
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
	err := pkg.FormValidate(c, &u)
	if err != nil {
		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST "+err.Error())
	}

	_, ok := h.store.FindOneByEmail(u.Email)
	if ok == nil {
		return c.String(http.StatusConflict, "ERR_CONFLICT user with given email already exists")
	}

	hashPwd, err := security.HashPassword(u.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER, can't hash password")
	}
	// generate hash password
	u.Password = hashPwd
	_, err = h.store.Create(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST"+err.Error())
	}
	//  generate JWT
	t, err := security.CreateAccesToken(u.Email)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER, can't create access token")
	}
	fmt.Println(t)
	return render(c, user.Show(u), layout.Base())
}
