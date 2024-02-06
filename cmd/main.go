package main

import (
	"context"

	"github.com/bilkadev/Go_HTMX_Real-chat/handler"
	"github.com/labstack/echo/v4"
)

func main() {
	app := echo.New()

	app.Static("/assets", "assets")
	userHandler := handler.UserHandler{}
	app.Use(withUser)

	app.GET("/user", userHandler.HandleUserShow)
	app.Start(":3000")
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "ctx@example.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
