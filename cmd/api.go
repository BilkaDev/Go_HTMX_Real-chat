package main

import (
	"context"

	"github.com/bilkadev/Go_HTMX_Real-chat/handler"
	"github.com/labstack/echo/v4"
)

type ApiServer struct {
	*echo.Echo
}

func NewApiServer() *ApiServer {
	return &ApiServer{}
}

func (*ApiServer) Run() {
	server := echo.New()
	_ = LoadEnv()

	server.Static("/assets", "assets")

	server.Use(withUser)

	SetupRoutes(server)

	server.Start(":3000")
}

func SetupRoutes(s *echo.Echo) {
	userHandler := handler.UserHandler{}
	s.GET("/user", userHandler.HandleUserShow)
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), "user", "ctx@example.com")
		c.SetRequest(c.Request().WithContext(ctx))
		return next(c)
	}
}
