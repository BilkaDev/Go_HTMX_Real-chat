package security

import (
	"fmt"
	"net/http"
	"time"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	Email string `json:"email"`
	jwt.RegisteredClaims
}

var tokenExpires = time.Hour * 24

const jwtKey = "jwt"

func CreateAccesToken(userName string) (string, error) {
	cfg := config.LoadEnv()
	claims := JwtCustomClaims{
		userName,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(tokenExpires)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(cfg.JWTSecret))

	if err != nil {
		fmt.Println(err)
		return t, err
	}

	return t, err
}

func WriteTokenCoocies(c echo.Context, token string) error {
	cookie := new(http.Cookie)
	cookie.Name = jwtKey
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.SameSite = 3
	cookie.Expires = time.Now().Add(tokenExpires)

	if ctx, ok := c.Get("env").(*config.Config); ok {
		if ctx.State != "dev" {
			cookie.Secure = true
		}
	}
	c.SetCookie(cookie)
	return nil
}

func ClearSeassion(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = jwtKey
	cookie.Value = ""
	cookie.MaxAge = 0
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	return nil
}
