package middleware

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/bilkadev/Go_HTMX_Real-chat/pkg/security"
	"github.com/labstack/echo/v4"
)

func CurrentUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		jwt, err := c.Cookie("jwt")
		if err != nil {
			return next(c)
		}
		t, err := security.VerifyAccessToken(jwt.Value)

		if err != nil {
			return next(c)
		}

		c.Set(config.CurrentUserId.String(), t.UserId)
		c.Set(config.CurrentUserName.String(), t.UserName)
		return next(c)
	}
}
