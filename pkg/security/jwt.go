package security

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/golang-jwt/jwt/v5"
	"github.com/labstack/echo/v4"
)

type JwtCustomClaims struct {
	UserName string `json:"UserName"`
	UserId   uint   `json:"UserId"`
	jwt.RegisteredClaims
}

var TOKEN_EXPIRES = time.Hour * 24
var JWT_SECRET string = os.Getenv("JWT_SECRET")
var JWT_KEY_COOKIES = "jwt"

func CreateAccesToken(userName string, userId uint) (string, error) {
	claims := JwtCustomClaims{
		userName,
		userId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TOKEN_EXPIRES)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(JWT_SECRET))

	if err != nil {
		fmt.Println(err)
		return t, err
	}

	return t, err
}

func VerifyAccessToken(tokenString string) (*JwtCustomClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JwtCustomClaims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte(JWT_SECRET), nil
	})
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*JwtCustomClaims); ok {
		return claims, nil
	} else {
		return nil, fmt.Errorf("unknown claims type, cannot proceed")
	}
}

func WriteTokenCoocies(c echo.Context, token string) error {
	cookie := new(http.Cookie)
	cookie.Name = JWT_KEY_COOKIES
	cookie.Value = token
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.SameSite = 3
	cookie.Expires = time.Now().Add(TOKEN_EXPIRES)

	if ctx, ok := c.Get("env").(*config.Env); ok {
		if ctx.State != "dev" {
			cookie.Secure = true
		}
	}
	c.SetCookie(cookie)
	return nil
}

func ClearSeassion(c echo.Context) error {
	cookie := new(http.Cookie)
	cookie.Name = JWT_KEY_COOKIES
	cookie.Value = ""
	cookie.HttpOnly = true
	cookie.Path = "/"
	cookie.SameSite = 3
	cookie.MaxAge = 0
	cookie.Expires = time.Now()
	c.SetCookie(cookie)
	return nil
}
