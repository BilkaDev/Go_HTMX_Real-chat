package main

import (
	"context"

	"github.com/bilkadev/Go_HTMX_Real-chat/handler"
	"github.com/bilkadev/Go_HTMX_Real-chat/types"
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
	config := LoadEnv()

	store := NewSqliteStore()

	server.Static("/assets", "assets")

	server.Use(withUser)

	SetupRoutes(server, store)

	server.Start(":" + config.Port)
}

func SetupRoutes(s *echo.Echo, storage *types.SqlStore) {
	handler.AuthRouter(s, "/auth", storage)
}

func withUser(next echo.HandlerFunc) echo.HandlerFunc {
	return SetContext(next, "user", "elo mordo")
}

func SetContext(n echo.HandlerFunc, key any, value any) echo.HandlerFunc {
	return func(c echo.Context) error {
		ctx := context.WithValue(c.Request().Context(), key, value)
		c.SetRequest(c.Request().WithContext(ctx))
		return n(c)
	}
}
