package main

import (
	"context"

	"github.com/bilkadev/Go_HTMX_Real-chat/config"
	"github.com/bilkadev/Go_HTMX_Real-chat/handler"
	"github.com/bilkadev/Go_HTMX_Real-chat/middleware"
	"github.com/bilkadev/Go_HTMX_Real-chat/store"
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
	config := config.LoadEnv()

	store := store.NewSqliteStore()

	server.Static("/assets", "assets")

	server.Use(middleware.LoggerMiddleware)

	server.Use(withUser)

	SetupRoutes(server, store)

	server.Start(":" + config.Port)
}

func SetupRoutes(s *echo.Echo, store *store.SqlStore) {
	handler.AuthRouter(s, "/auth", store)
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
