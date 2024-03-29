package main

import (
	"net/http"

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
	env := config.LoadEnv()

	store := store.NewSqliteStore()
	server.Static("/", "public")

	server.Use(SetEnv(env))
	server.Use(middleware.CurrentUser)
	server.Use(middleware.LoggerMiddleware)

	SetupRoutes(server, store)

	server.Start(":" + env.Port)
}

func SetupRoutes(s *echo.Echo, store *store.SqlStore) {
	g := s.Group("/api/v1")
	g.GET("/health", func(c echo.Context) error {
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "OK"})
	})
	handler.HomeRouter(g, "/home", store)
	handler.AuthRouter(g, "/auth", store)
	handler.UserRouter(g, "/user", store)
	handler.MessageRouter(g, "/message", store)
}

func SetContext(n echo.HandlerFunc, key string, value any) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Set(key, value)
		return n(c)
	}
}

func SetEnv(env *config.Env) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return SetContext(next, config.EnvKey.String(), env)
	}
}
