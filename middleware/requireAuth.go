package middleware

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func RequireAuth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		user := c.Get("currentUserName")
		if user == nil {
			c.Response().Header().Set("hx-redirect", "/login")
			return c.String(http.StatusUnauthorized, "ERR_UNAUTHORIZED ")
		}
		return next(c)
	}
}
