package handler

import (
	"net/http"

	"github.com/a-h/templ"
	"github.com/bilkadev/Go_HTMX_Real-chat/model"
	"github.com/bilkadev/Go_HTMX_Real-chat/pkg"
	"github.com/bilkadev/Go_HTMX_Real-chat/pkg/security"
	"github.com/bilkadev/Go_HTMX_Real-chat/store"
	"github.com/bilkadev/Go_HTMX_Real-chat/view/home"
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
	g.POST("/login", ah.HandleAuthLogin)
	g.POST("/signup", ah.HandleAuthSignUp)
	g.POST("/logout", ah.HandleAuthLogout)
}

func (h AuthHandler) HandleAuthLogin(c echo.Context) error {
	u := model.UserIn{}
	err := pkg.FormValidate(c, &u)
	if err != nil {
		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST "+err.Error())
	}

	us, err := h.store.FindOneByUserName(u.UserName)
	if err != nil {
		return c.String(http.StatusUnauthorized, "ERR_UNAUTHORIZED Invalid userName or password")
	}

	valid := security.CheckPasswordHash(u.Password, us.Password)
	if !valid {
		return c.String(http.StatusUnauthorized, "ERR_UNAUTHORIZED Invalid userName or password")
	}
	t, err := security.CreateAccesToken(u.UserName, us.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER, can't create access token")
	}

	security.WriteTokenCoocies(c, t)

	c.Response().Header().Set("HX-Push-Url", "/")
	return render(c, home.Show([]model.User{}, templ.Attributes{}))
}

func (h AuthHandler) HandleAuthLogout(c echo.Context) error {
	//  delete coockies from session
	security.ClearSeassion(c)
	c.Response().Header().Set("hx-redirect", "/login")
	return c.String(http.StatusNoContent, "logged out sccessfuly")
}

func (h AuthHandler) HandleAuthSignUp(c echo.Context) error {
	u := model.User{}
	err := pkg.FormValidate(c, &u)
	if err != nil {
		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST "+err.Error())
	}
	if u.Password != c.FormValue("repassword") {

		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST password dosen't match")
	}

	_, ok := h.store.FindOneByEmail(u.Email)
	if ok == nil {
		return c.String(http.StatusConflict, "ERR_CONFLICT user with given email already exists")
	}
	_, ok = h.store.FindOneByUserName(u.UserName)
	if ok == nil {
		return c.String(http.StatusConflict, "ERR_CONFLICT user with given userName already exists")
	}

	hashPwd, err := security.HashPassword(u.Password)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER, can't hash password")
	}
	// generate hash password
	u.Password = hashPwd

	// random avatar pics

	g := "boy"
	if u.Gender == "female" {
		g = "girl"
	}
	u.Avatar = "https://avatar.iran.liara.run/public/" + g + "?username=" + u.UserName

	_, err = h.store.Create(&u)
	if err != nil {
		return c.String(http.StatusBadRequest, "ERR_BAD_REQUEST"+err.Error())
	}
	//  generate JWT
	t, err := security.CreateAccesToken(u.UserName, u.ID)
	if err != nil {
		return c.String(http.StatusInternalServerError, "ERR_INTERNAL_SERVER, can't create access token")
	}

	security.WriteTokenCoocies(c, t)
	c.Response().Header().Set("HX-Push-Url", "/")
	return render(c, home.Show([]model.User{}, templ.Attributes{}))
}
