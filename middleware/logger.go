package middleware

import (
	"github.com/bilkadev/Go_HTMX_Real-chat/pkg"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoMiddleware "github.com/labstack/echo/v4/middleware"
	"go.uber.org/zap"
)

func LoggerMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	logger := pkg.NewLogger()
	return echoMiddleware.RequestLoggerWithConfig(echoMiddleware.RequestLoggerConfig{
		LogURI:    true,
		LogStatus: true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("request",
				zap.String("URI", v.URI),
				zap.Int("status", v.Status),
			)

			return nil
		},
	})(next)
}
